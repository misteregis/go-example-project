# Documentação Swagger - Guia de Uso

## 📚 Visão Geral

Este projeto utiliza **Swagger/OpenAPI** para documentação automática da API. A documentação é gerada a partir de anotações nos comentários do código Go.

## 🚀 Acesso Rápido

- **URL Principal:** http://localhost:8080/ (redireciona automaticamente)
- **Swagger UI:** http://localhost:8080/swagger/index.html
- **Swagger JSON:** http://localhost:8080/swagger/doc.json

## 📝 Como Adicionar uma Nova Rota

### Passo 1: Criar o Handler com Anotações

```go
// GetProducts godoc
// @Summary Listar todos os produtos
// @Description Retorna a lista de todos os produtos disponíveis
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {string} string "Erro interno"
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
  // Implementação
}
```

### Passo 2: Registrar a Rota no main.go

```go
api.HandleFunc("/products", handlers.GetProducts).Methods("GET")
```

### Passo 3: Regenerar a Documentação

```bash
make swagger
# ou
swag init -g cmd/api/main.go -o docs
```

### Passo 4: Reiniciar o Servidor

```bash
make run
```

A nova rota aparecerá automaticamente no Swagger UI!

## 🏷️ Anotações Swagger

### Anotações Principais

- **@Summary** - Breve descrição do endpoint
- **@Description** - Descrição detalhada
- **@Tags** - Agrupa endpoints relacionados
- **@Accept** - Tipo de conteúdo aceito (json, xml, etc)
- **@Produce** - Tipo de conteúdo retornado
- **@Param** - Parâmetros da requisição
- **@Success** - Resposta de sucesso
- **@Failure** - Respostas de erro
- **@Router** - Caminho e método HTTP

### Tipos de Parâmetros

```go
// Parâmetro de Path
// @Param id path int true "ID do Usuário"

// Parâmetro de Query
// @Param name query string false "Nome para filtrar"

// Parâmetro de Body
// @Param user body models.User true "Dados do Usuário"

// Parâmetro de Header
// @Param Authorization header string true "Token de Autenticação"
```

### Respostas

```go
// Resposta com objeto
// @Success 200 {object} models.User

// Resposta com array
// @Success 200 {array} models.User

// Resposta simples
// @Success 200 {string} string "OK"

// Múltiplas respostas de erro
// @Failure 400 {string} string "Requisição inválida"
// @Failure 404 {string} string "Não encontrado"
// @Failure 500 {string} string "Erro interno"
```

## 📦 Estrutura de Models

Para documentar melhor os schemas, use tags `example` nos models:

```go
type Product struct {
  ID          int     `json:"id" example:"1"`
  Name        string  `json:"name" example:"Produto Exemplo"`
  Price       float64 `json:"price" example:"99.99"`
  Description string  `json:"description" example:"Descrição do produto"`
}
```

## 🔧 Configuração Global

As configurações globais estão no `main.go`:

```go
// @title API de Exemplo em Go
// @version 1.0
// @description API RESTful de exemplo com documentação Swagger automática

// @contact.name Suporte API
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
```

## 📖 Exemplo Completo: CRUD de Produtos

### 1. Criar o Model

```go
// internal/models/product.go
package models

type Product struct {
  ID          int     `json:"id" example:"1"`
  Name        string  `json:"name" example:"Notebook"`
  Price       float64 `json:"price" example:"2999.99"`
  Description string  `json:"description" example:"Notebook de alta performance"`
}
```

### 2. Criar os Handlers

```go
// internal/handlers/product.go
package handlers

// GetProducts godoc
// @Summary Listar produtos
// @Description Retorna todos os produtos
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
  // Implementação
}

// CreateProduct godoc
// @Summary Criar produto
// @Description Cria um novo produto
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Dados do Produto"
// @Success 201 {object} models.Product
// @Failure 400 {string} string "Dados inválidos"
// @Router /products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
  // Implementação
}
```

### 3. Registrar as Rotas

```go
// cmd/api/main.go
api.HandleFunc("/products", handlers.GetProducts).Methods("GET")
api.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
```

### 4. Regenerar Documentação

```bash
make swagger
```

## 🎯 Boas Práticas

1. **Sempre adicione anotações** ao criar novos endpoints
2. **Use tags descritivas** para agrupar endpoints relacionados
3. **Documente todos os parâmetros** incluindo tipo e obrigatoriedade
4. **Inclua exemplos** nos models usando a tag `example`
5. **Documente erros comuns** com @Failure
6. **Regenere a documentação** após mudanças: `make swagger`

## 🔄 Workflow Recomendado

1. Criar/modificar handler com anotações Swagger
2. Registrar rota no main.go
3. Executar `make swagger`
4. Executar `make run`
5. Testar no Swagger UI (http://localhost:8080/)

## 📚 Recursos Adicionais

- [Documentação Oficial Swaggo](https://github.com/swaggo/swag)
- [Especificação OpenAPI](https://swagger.io/specification/)
- [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format)

## ⚡ Comandos Úteis

```bash
# Regenerar documentação
make swagger

# Instalar/atualizar swag
go install github.com/swaggo/swag/cmd/swag@latest

# Validar documentação
swag init -g cmd/api/main.go -o docs --parseVendor

# Limpar e regenerar
rm -rf docs && make swagger
```

## 🎨 Customização

Para personalizar o Swagger UI, você pode modificar as configurações em `cmd/api/main.go`:

```go
// Configuração customizada
r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
  httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
  httpSwagger.DeepLinking(true),
  httpSwagger.DocExpansion("none"),
  httpSwagger.DomID("swagger-ui"),
))
```
