# format whole project with go fmt
format : 
	go fmt ./...

# Runs all test with clean test cache and also shows test coverage for all files
test : 
	go clean -testcache && go test -v -cover ./...

# Shows project's test coverage
coverage :
	go test -cover ./...

# Run on local
run:
	go run main.go