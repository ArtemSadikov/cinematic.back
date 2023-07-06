proto:
	protoc -I=./proto \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./internal/delivery/grpc/interface \
		--go_out=paths=source_relative:./internal/delivery/grpc/interface \
	proto/auth.proto

.PHONY: proto