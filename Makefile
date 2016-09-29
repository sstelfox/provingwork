run:
	@go run examples/hashcash_example.go
	@go run examples/strongwork_example.go

setup:
	go get github.com/tools/godep
	godep restore

.DEFAULT_GOAL := run
.PHONY: run setup
