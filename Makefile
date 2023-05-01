GO=GOFLAGS=-mod=vendor go

.PHONY: run
run:
	$(GO) run src/main.go

.PHONY: deps
deps:
	$(GO) mod tidy