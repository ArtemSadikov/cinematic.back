path = ./internal/delivery/grpc/interface

proto_user_model:
	protoc -I=./proto -I=./proto/users \
		--go_out=paths=source_relative:$(path) \
	./proto/users/users_model.proto

proto_user_service:
	protoc -I=./proto -I=./proto/users \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:$(path) \
		--go_out=paths=source_relative:$(path) \
	./proto/users/users_service.proto

proto_auth_model:
	protoc -I=./proto -I=./proto/auth \
		--go_out=paths=source_relative:$(path) \
	proto/auth/auth_model.proto

proto_auth_service:
	protoc -I=./proto -I=./proto/auth \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:$(path) \
		--go_out=paths=source_relative:$(path) \
	proto/auth/auth_service.proto

proto_all: proto_user_model proto_user_service proto_auth_model proto_auth_service

create_proto:
	$(MAKE) proto_all path=./internal/delivery/grpc/interface
	$(MAKE) proto_all path=../../pkg/clients/users/interface

build_api:
	go build -o build/bin ./cmd/api/main.go