## Network Block Store

Network Block Device implementation over YDB BlobStorage or over our own storage nodes. Offers reliable thin-provisioned block devices which support snapshots.

Block storage [overview diagram](https://github.com/ydb-platform/nbs/blob/main/doc/blockstore/overview/overview.png?raw=true)

### Quickstart

Follow the instructions [here](VSCODE.md) to generate workspace and install the necessary plugins.

Follow the instructions [here](example/README.md) to build and run NBS on your machine and to attach an NBS-based disk via NBD. NBS-based disks can be attached via vhost-user-blk as well.

Follow the instructions [here](CLANG-FORMAT.md) to install clang-format for formatting the code.

Additional information about features of our [Github Actions](GITHUB.md) (labels, test results and so on)

### Documentation

The docs can be found [here](/doc). We are in the process of writing them atm. The overall repository structure can be found [here](/doc/REPOSITORY_STRUCTURE.md).

### How to Deploy

TODO
