api_bin_path=build/api/bin

generate_schema:
	cd internal/api/graphql && go run github.com/99designs/gqlgen generate

build_api_server:
	go build -o $(api_bin_path) cmd/api/main.go

run_api_server:
	$(api_bin_path)

run_server:
ifeq ("$(wildcard $(api_bin_path))","")
	$(MAKE) build_api_server
endif
	$(MAKE) run_api_server

.PHONY: generate_schema run_server