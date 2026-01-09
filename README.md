# Go Example Project

Um projeto de exemplo em Go demonstrando boas práticas de desenvolvimento.

## 📋 Pré-requisitos

- Go 1.21 ou superior
- VSCode (recomendado)
- Docker (opcional)

## 🚀 Início Rápido

### Instalação

```bash
# Clone o repositório
git clone https://github.com/misteregis/go-example-project. git
cd go-example-project

# Instale as dependências
make deps

# Instale as ferramentas de desenvolvimento
make install-tools
```

### Executar a aplicação

```bash
make run
```

A aplicação estará disponível em `http://localhost:8080`

## 🛠️ Comandos Disponíveis

```bash
make help           # Mostra todos os comandos disponíveis
make run            # Executa a aplicação
make build          # Compila a aplicação
make test           # Executa os testes
make test-coverage  # Mostra cobertura de testes
make lint           # Executa o linter
make fmt            # Formata o código
make clean          # Remove arquivos compilados
```

## 📁 Estrutura do Projeto

```
.
├── cmd/            # Pontos de entrada da aplicação
├── internal/       # Código privado da aplicação
├── pkg/            # Bibliotecas públicas reutilizáveis
├── .vscode/        # Configurações do VSCode
└── Makefile        # Comandos de automação
```

## 🔧 Configuração do VSCode

O projeto inclui configurações recomendadas do VSCode em `.vscode/`:
- **extensions.json**: Extensões recomendadas
- **settings. json**: Configurações de formatação e linting

Ao abrir o projeto, o VSCode sugerirá a instalação das extensões recomendadas.

## 📝 Endpoints da API

### Documentação Swagger

A API possui documentação interativa Swagger disponível em:
- **Acesse:** `http://localhost:8080/` (redireciona automaticamente para o Swagger)
- **Ou diretamente:** `http://localhost:8080/swagger/index.html`

### Regenerar Documentação

Sempre que adicionar novas rotas ou modificar as anotações, execute:

```bash
swag init -g cmd/api/main.go -o docs
```

### Endpoints Disponíveis

- `GET /` - Redireciona para /swagger/index.html
- `GET /swagger/` - Documentação Swagger UI
- `GET /api/v1/health` - Health check
- `GET /api/v1/users` - Lista todos os usuários
- `GET /api/v1/users/{id}` - Obtém um usuário específico
- `POST /api/v1/users` - Cria um novo usuário

## 🧪 Testes

```bash
# Executar todos os testes
make test

# Ver cobertura de testes
make test-coverage
```

## 📦 Build

```bash
# Compilar a aplicação
make build

# O binário será gerado em ./bin/api
./bin/api
```

## 🎨 Linting e Formatação

O projeto usa **golangci-lint** para análise estática de código:

```bash
# Executar o linter
make lint

# Formatar código
make fmt
```

## 📄 Licença

MIT License
