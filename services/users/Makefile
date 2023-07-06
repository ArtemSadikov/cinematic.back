proto_model:
	protoc -I=./proto \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./internal/delivery/grpc/interface \
		--go_out=paths=source_relative:./internal/delivery/grpc/interface \
	proto/users_model.proto

proto_service:
	protoc -I=./proto \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./internal/delivery/grpc/interface \
		--go_out=paths=source_relative:./internal/delivery/grpc/interface \
	proto/users_service.proto


proto_all: proto_model proto_service

build_api:
	go build -o build/bin ./cmd/api/main.go