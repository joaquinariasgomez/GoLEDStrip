GO=go

.PHONY: run
run:
	$(GO) run src/main.go

.PHONY: deps
deps:
	$(GO) mod tidy

# Para ejecutar: sudo go run src/main.go