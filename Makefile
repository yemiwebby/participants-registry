server:
	go run main.go
test:
	GIN_MODE=release gotest -v -cover ./api