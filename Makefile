unittests:
	@go clean -testcache
	@go test ./...

build:
	@cd gogpt; go build .; cd -

clean:
	@rm -f gogpt/gogpt