run:
	@go run examples/hashcash_example.go
	@go run examples/strongwork_example.go

.DEFAULT_GOAL := run
.PHONY: run
