run:
	go run ./cmd/users/main.go --config=./config/dev.yaml

mod:
	go mod tidy && go mod vendor

generate:
	mkdir -p ./pkg/users_v1
	protoc --go_out=./pkg/users_v1 \
	--go_opt=paths=source_relative \
	--go-grpc_out=./pkg/users_v1 \
	--go-grpc_opt=paths=source_relative ./api/users_v1/users.proto
