# C2 Chopper Framework

### Prerequisites

* Protocol buffer compiler, `protoc`. [Installation Guide.](https://grpc.io/docs/protoc-installation/)
* Go plugins for the protocol compiler.
    1. Install the protocol compiler plugins for Go using the following commands:
    ```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```

    1. Install gRPC-Gateway plugin:
    ```
    $ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
    $ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
    ```

    1. Update your PATH so that the protoc compiler can find the plugins:
    ```
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```

## HTTP -> GRPC Gateway
