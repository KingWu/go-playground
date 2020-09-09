# Run make command by following
# make <command>
build:
	go build
format:
	go fmt
run:
	go run main.go
test:
	go test ./...
gqlgen:
	go run github.com/99designs/gqlgen generate
