export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"

for file in $(find ./proto/ -name "*.proto"); do
       protoc --go_out=paths=source_relative:./proto \
              --go-grpc_out=paths=source_relative:./proto \
              --proto_path=./proto/ \
              $file
done
