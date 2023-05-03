.PYTHON: lint-all, lint, vet
lint-all: lint vet

lint:
	staticcheck ./...

vet:
	go vet ./...