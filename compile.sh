export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"

for file in $(find ./grpc/proto-files -name "*.proto"); do
	protoc -I grpc/proto/google/api \
        --go_out=paths=source_relative:./grpc/proto \
		--go-grpc_out=paths=source_relative:./grpc/proto \
		--proto_path=./grpc/proto-files \
        --grpc-gateway_out ./grpc/proto \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
		$file





	# protoc --go-grpc_out=Mclient:client/grpc/ \
	# 	--go-grpc_out=Mserver:server/grpc/ \
	# 	--proto_path=./proto/ \
	# 	$file
done
