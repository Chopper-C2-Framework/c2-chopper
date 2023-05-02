export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"

protoc --proto_path=./proto/ \
       --go_out=paths=source_relative:./proto \
       --go-grpc_out=paths=source_relative:./proto \
       proto/*.proto
