
# dfsmr

`dfsmr` or Distributed Finite State Machine Runner is an experiment in interface design
for what an expressive, modern distributed task management system might look like. Worth
repeating again, this is an *experiment* in *interface design*. The implementation works,
but only to extent necessary to exercise the interfaces, and is no way performant, reliable,
etc.

Most modern task systems such as RabbitMQ or Celery focus on queuing tasks and having a
distributed fleet of consumers pull down tasks. This works remarkably well, but also offloads
expressing task workflow to the applications themselves. If you want sophisticated workflows,
you end up in a parallel universe of tools like Airflow, which are excellent but heavy and designed
to coordinate batch workflows rather than task workflows.

`dfsmr` envisions a world where you could have performant, expressive task workflows, represented
as a fleet of finite state machines, similar to [Erlang's gen_fsm](http://erlang.org/doc/man/gen_fsm.html),
but with instances and transitions coordinated and enforced by a centralized system.

You start by defining a *machine*, which defines any number of nodes and transitions that comprise your
state machine. Machine definitions are written in YAML, for example:

```
id: crawler
nodes:
  crawl:
    start: true
    transitions:
      ok: success
      error: wait
  wait:
    transitions:
      ok: crawl
  success:
    final: true
```

Once you've defined the machine, you load it via the `Define` endpoint
or the `dcli` command line tool:

```
dcli define ./crawl.fsm.yaml
```

Once defined, you can list all machines via:

```
dcli machines
```

Next, you'll want to create an *instance* of the *machine*, which is instantiating the
machine with a set of input parameters. This is done via the `Start` endpoint.
Each instance has a unique identifier, is passed zero or more key-value pairs as initial inputs.

```
dcli start crawler
# TODO: should be able to pass key-value pairs via dcli, although
# if you call the grpc endpoint directly, you can indeed do this!
``

Once the instance is created, you can list all instances and their states via:

```
dcli instances
```

At this point consumers are able to retrieve instances and perform work on them.
As a consumer, you can request any available instance, or you can choose to filter
by machine or current node. For example if you wanted a dedicated crawler consumer
that only performed crawl operations, you'd do so via

```
dcli ready crawler crawl
```

If you wanted any instance ready for work, you'd simply:

```
dcli ready
```

Once you've performed work for a given node, you update the node by
signaling a transition via the `Transition` endpoint. After each transition,
you can supply a new set of key-value pairs for the next node.

Somewhat interestingly, a consumer performing the work doesn't need to be
aware of the next node in the state machine, although it does need to be
aware of the transition. For example, you might signal the `ok` transition,
which `dfsm` applies to the machine definition to move the instance into
the `success` node, but you could seamlessly insert another validation step
in without requiring the consumer to update its logic. (There is some complexity
here around ensuring the input key-value pairs don't leak state transition logic,
long term I think you'd want to enforce a single protobuf definition that must be
used across every transition in a state machine, but I decided to punt on that
given this is only an experiment. :)

## Running server

Building the server is:

    dep ensure
    make server

Running the server is:

    ./dsrv
    ./dsrv --addr :3004

## Using client

Building the client is as simple as:

    dep ensure
    make client

Using the client is:

    $ ./dcli start
    2018/04/29 08:40:51 Started success:true successMessage:"Success"

    $ ./dsrv --addr :3004 start
    2018/04/29 08:40:51 Started success:true successMessage:"Success"    

    $ ./dcli changes

The `Changes` endpoint and command are inspired by [Redis' MONITOR command](https://redis.io/commands/monitor)
and is useful for debugging / understanding what's happening.


## Commands

[See the grpc definition for all commands.](./blob/master/dfsmr/dfsmr.proto)

# Unimplemented ideas for extension

This section discusses some ideas for implementing some interesting aspects
of what I think a production version of this would need to address.

## Handling instance failure: timeouts and retries

There are a number of ways that processing a given instance can fail.
Generally our design philosophy is to provide as much customization
as possible within the machine definitions, so we don't provide an
implicit retry mechanism, rather you'd add a `retry` node:

```
crawl -[ok]-> crawled
crawl -[error]-> retry -> -[ok]-> crawl
```

It's a bit toilsome to write, but it allows you to implement arbitrary
sophisticated retry logic without requiring changes to the platform itself.

Retries are different, as timing out is a transition caused implicitly by inaction,
as opposed to a response to an explicit transition. As such, timeouts are a core
feature provided by `dfsmr`.

By default, timeouts are dynamically generated for each node, using the
p95 duration of previous attempts. When a timeout is fired, the instance
becomes available to schedule on other workers.

```
nodes:
  validate:
    # these are default values
    timeout_strategy: percentile
    # timeout percentile
    timeout: 95.0
```

You can also specify fixed timeout values, although generally
this is not recommended as it becomes toilsome to maintain these
values as your system evolves:

```
nodes:
  validate:
    timeout_strategy: fixed
    # timeout in seconds
    timeout: 30.0
```

Because we're using the p95, as opposed to e.g. the p99.9, it is relatively
common to have multiple workers performing work on behalf of the same instance,
which is an intentional choice to ensure you design for that scenario rather
than allow emergent behavior to reign.


When multiple workers are performing work on behalf of the same instance,
the first one to complete will be accepted, and any subsequent completions
of the same transition will be rejected as invalid transitions. The one
caveat to this is related to *Versions*, described later--depending on
the upgrade strategy the machine has defined--an upgraded machine
may reject work started during a previous version.

There is a circuit-breaker on retries for each machine. If the retry
rate becomes extreme, retries will be severely throttled until they begin
to succeed. (The most common case here would be a bad deploy in your workers,
or a downstream dependency of the workers degrading.)

In cases where your transitions are not-idempotent and are dangerous to
timeout, you can explicitly set that node's transition timeout value to `0.0`,
which signals to never retry

```
nodes:
  validate:
    # value in seconds
    timeout: 0.0
```

This will cause the behavior for a given node to switch from "at least once" to
"at most once". If a failure occurs, you'll have to handle it explicitly yourself.
This is a significant anti-pattern, and if you're needing to use it, then it implies
a potentially significant architecture flaw in your application, that will grow
increasingly difficult to manage as your system scales functionality and load.

## Versions

One of the challenges of a distributed task runner is providing
expressive semantics for handling change, particularly as it relates
to changes to a FSM. For example, if you updated your crawler FSM,
you might prefer the system to abort existing instances, but for
very long-running jobs, you might prefer to complete current jobs
according to the machine definition that it started under.

We allow each machine to select the behavior it wants by specifying
a `upgrade_strategy` configuration within the machine's configuration YAML:

```
# continue finishes running instances using the configuration they
# were launched under
# this is the default value
upgrade_strategy: continue

# upgrade: restart running fsms under new configuration
upgrade_strategy: upgrade

# abort causes all jobs started previously to be failed
upgrade_strategy: abort
```

, through
a versioning mechanism.


## System node and transitions

In addition to the nodels and transitions defined within your machine config,
there are also these system nodes added to every machine:

- `dfsmr.upgraded` is used during machine config changes, when upgrading
    an already running instance to a new version
- `dfsmr.aborted` is used during machine config changes, when aborting already running instances
- `dfsmr.timeout` is used when a transition has timed out and made available
    for reprocessing


