run:
	go run ./cmd/users/main.go --config=./config/dev.yaml

mod:
	go mod tidy && go mod vendor
