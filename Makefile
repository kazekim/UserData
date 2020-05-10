prepare:
	go get ./...
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/userdata-linux-x64 cmd/userdata/main.go
run:
	go run cmd/userdata/main.go