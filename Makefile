test:
	go test -v -p 1 -cover -failfast ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'

test-cover: test
	go tool cover -html=coverage.out