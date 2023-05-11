export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"

for file in $(find ./grpc/proto -name "*.proto"); do
	protoc --go_out=paths=source_relative:./grpc/interfaces \
		--go-grpc_out=paths=source_relative:./grpc/interfaces \
		--proto_path=./grpc/proto \
		$file

	# protoc --go-grpc_out=Mclient:client/grpc/ \
	# 	--go-grpc_out=Mserver:server/grpc/ \
	# 	--proto_path=./proto/ \
	# 	$file
done
