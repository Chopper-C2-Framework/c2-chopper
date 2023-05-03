# C2 Chopper Framework

### Prerequisites

* Protocol buffer compiler, `protoc`. [Installation Guide.](https://grpc.io/docs/protoc-installation/)
* Go plugins for the protocol compiler. 
    1. Install the protocol compiler plugins for Go using the following commands:
    ```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```

    2. Update your PATH so that the protoc compiler can find the plugins:
    ```
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```