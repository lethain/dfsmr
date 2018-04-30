

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
