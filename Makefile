unittests:
	@go clean -testcache
	@go test ./...

codecoverage:
	@go test -coverprofile cover.out ./...
	@go tool cover -html=cover.out

build:
	@cd gogpt; go build .; cd -

clean:
	@rm -f gogpt/gogpt