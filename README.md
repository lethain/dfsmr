

Vocabulary:
- Machine: a defined finite state machine, including its nodes, transitions, inputs and configuration
- Instance: is a specific invocation of a Machine, triggered by `start` command
- Worker: external process that calls the `ready` command to request work

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


# Running `dfsmr`

two modes:
1. standalone local mode, uses leveldb for storage, no other dependencies
2. distributed mode, uses Kafka



## Operations

* Start: machine, id, context
* Transition: id, transition, input
* Relinquish: id
* Ready: machine, allowed_transitions, max_batch_size

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
