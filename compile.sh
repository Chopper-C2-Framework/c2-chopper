export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"

for file in $(find ./grpc/proto-files -name "*.proto"); do
	protoc --go_out=paths=source_relative:./grpc/proto \
		--go-grpc_out=paths=source_relative:./grpc/proto \
		--proto_path=./grpc/proto-files \
		$file

	# protoc --go-grpc_out=Mclient:client/grpc/ \
	# 	--go-grpc_out=Mserver:server/grpc/ \
	# 	--proto_path=./proto/ \
	# 	$file
done
