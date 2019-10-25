unittests:
	@go test ./...

build:
	@cd gogpt; go build .; cd -

clean:
	@rm -f gogpt/gogpt