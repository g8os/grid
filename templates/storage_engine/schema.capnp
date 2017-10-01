@0x80bf2ccb5d7de793;

struct Schema {
    homeDir @0 :Text; # directory where the storageEngine db will be stored
    bind @1: Text; # listen bind address.

    master @2 :Text;
    # name of other storageEngine service that needs to be used as master
    # if this is filled, this instance will behave as a slave

    container @3 :Text; # pointer to the parent service
    status @4: Status;

    enum Status{
        halted @0;
        halting @1;
        running @2;
        unhealthy @3;
    }
}
