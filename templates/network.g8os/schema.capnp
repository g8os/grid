@0xeee35f60471101a3;

struct Schema {
    vlanTag @0 :UInt16;
    cidr @1 :Text; # of the storage network
    vxaddr @2 :Text;
    storageaddr @3 :Text;
    subnet @4 :Text;
}
