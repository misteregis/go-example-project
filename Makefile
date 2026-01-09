.PHONY: help run build test lint clean install-tools deps tidy fmt test-coverage docker-build swagger

help:  ## Mostra esta mensagem de ajuda
	@echo Comandos disponiveis:
	@echo   make install-tools - Instala ferramentas de desenvolvimento
	@echo   make swagger       - Gerar documentacao Swagger
	@echo   make run           - Executa a aplicacao
	@echo   make build         - Compila a aplicacao
	@echo   make test          - Executa os testes
	@echo   make test-coverage - Mostra cobertura de testes
	@echo   make lint          - Executa o linter
	@echo   make fmt           - Formata o codigo
	@echo   make clean         - Remove arquivos compilados
	@echo   make deps          - Baixa dependencias
	@echo   make tidy          - Organiza dependencias

install-tools: ## Instala ferramentas de desenvolvimento
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/swaggo/swag/cmd/swag@latest

swagger: ## Gera documentacao Swagger
	swag init -g cmd/api/main.go -o docs

run: ## Executa a aplicacao
	go run cmd/api/main.go

build: ## Compila a aplicacao
	go build -o bin/api.exe cmd/api/main.go

test: ## Executa os testes
	go test -v -coverprofile=coverage.out ./...

test-coverage: test ## Mostra cobertura de testes
	go tool cover -html=coverage.out

lint: ## Executa o linter
	golangci-lint run ./...

fmt: ## Formata o codigo
	gofmt -s -w .
	goimports -w .

clean: ## Remove arquivos compilados
	@if exist bin rmdir /s /q bin
	@if exist coverage.out del /f coverage.out

deps: ## Baixa dependencias
	go mod download
	go mod verify

tidy: ## Organiza dependencias
	go mod tidy

docker-build: ## Cria imagem Docker
	docker build -t go-example-project:latest .

.DEFAULT_GOAL := help
