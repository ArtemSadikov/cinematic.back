.PHONY: proto

proto:
	@if [ ! -d "api/$(service)/pb" ]; then echo "creating new proto files..." && mkdir api/$(service)/pb; fi
	$(foreach proto_file, \
		$(shell find api/$(service)/proto -name '*.proto'), \
			protoc -I=./api/$(service)/proto \
				--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:api/$(service)/pb \
				--go_out=paths=source_relative:api/$(service)/pb \
		$(proto_file);)