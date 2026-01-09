# Testando a API com Swagger

## 🚀 Como Usar o Swagger UI

### 1. Iniciar o Servidor

```bash
make run
```

### 2. Acessar o Swagger UI

Abra o navegador em: http://localhost:8080/

Você será automaticamente redirecionado para: http://localhost:8080/swagger/index.html

### 3. Testar Endpoints

#### GET /api/v1/health

1. Clique em **"GET /api/v1/health"**
2. Clique em **"Try it out"**
3. Clique em **"Execute"**
4. Veja a resposta abaixo

**Resposta esperada:**
```json
{
  "status": "OK",
  "message": "Servidor funcionando corretamente"
}
```

#### GET /api/v1/users

1. Clique em **"GET /api/v1/users"**
2. Clique em **"Try it out"**
3. Clique em **"Execute"**

**Resposta esperada:**
```json
[
  {
    "id": 1,
    "name": "João Silva",
    "email": "joao@example.com",
    "age": 30,
    "created_at": "2026-01-09T10:00:00Z",
    "updated_at": "2026-01-09T10:00:00Z"
  },
  {
    "id": 2,
    "name": "Maria Santos",
    "email": "maria@example.com",
    "age": 25,
    "created_at": "2026-01-09T10:00:00Z",
    "updated_at": "2026-01-09T10:00:00Z"
  }
]
```

#### GET /api/v1/users/{id}

1. Clique em **"GET /api/v1/users/{id}"**
2. Clique em **"Try it out"**
3. No campo **"id"**, digite: `1`
4. Clique em **"Execute"**

**Resposta esperada:**
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@example.com",
  "age": 30,
  "created_at": "2026-01-09T10:00:00Z",
  "updated_at": "2026-01-09T10:00:00Z"
}
```

#### POST /api/v1/users

1. Clique em **"POST /api/v1/users"**
2. Clique em **"Try it out"**
3. Edite o JSON no campo **"Request body"**:

```json
{
  "name": "Pedro Oliveira",
  "email": "pedro@example.com",
  "age": 28
}
```

4. Clique em **"Execute"**

**Resposta esperada (Status 201 Created):**
```json
{
  "id": 3,
  "name": "Pedro Oliveira",
  "email": "pedro@example.com",
  "age": 28,
  "created_at": "2026-01-09T11:30:00Z",
  "updated_at": "2026-01-09T11:30:00Z"
}
```

## 📊 Schemas

O Swagger exibe automaticamente os schemas dos modelos. Clique em **"Schemas"** no final da página para ver:

- **models.User** - Estrutura completa do usuário com exemplos

## 🎯 Recursos do Swagger UI

### Testar Diferentes Respostas

O Swagger mostra todas as possíveis respostas:
- ✅ **200** - Sucesso
- ⚠️ **400** - Requisição inválida
- ❌ **404** - Não encontrado
- ❌ **500** - Erro interno

### Copiar cURL

Depois de executar uma requisição, você pode:
1. Rolar até **"Curl"**
2. Copiar o comando cURL
3. Executar no terminal

Exemplo:
```bash
curl -X 'GET' \
  'http://localhost:8080/api/v1/users' \
  -H 'accept: application/json'
```

### Ver Request/Response

O Swagger mostra:
- **Request URL** - URL completa da requisição
- **Server response** - Corpo da resposta
- **Response headers** - Cabeçalhos HTTP
- **Response code** - Código de status HTTP

## 🔧 Testando Validações

### Teste 1: Criar usuário sem nome

```json
{
  "email": "teste@example.com",
  "age": 25
}
```

**Resultado esperado:** Erro 400 - "nome inválido"

### Teste 2: Criar usuário sem email

```json
{
  "name": "Teste Silva",
  "age": 25
}
```

**Resultado esperado:** Erro 400 - "email inválido"

### Teste 3: Buscar usuário inexistente

GET /api/v1/users/999

**Resultado esperado:** Erro 404 - "Usuário não encontrado"

## 🎨 Dicas Úteis

1. **Favoritar no navegador:** http://localhost:8080/
2. **Recarregar a página** após regenerar o Swagger (`make swagger`)
3. **Usar "Try it out"** para testar rapidamente sem precisar de Postman/Insomnia
4. **Exportar** a definição OpenAPI (Download swagger.json) para usar em outras ferramentas
5. **Compartilhar** o link do Swagger com outros desenvolvedores

## 📱 Outras Formas de Acessar a Documentação

### JSON Raw
http://localhost:8080/swagger/doc.json

### YAML Raw
Verificar em: `docs/swagger.yaml`

### Importar no Postman/Insomnia

1. Copie o conteúdo de `docs/swagger.json`
2. No Postman/Insomnia, vá em **Import**
3. Cole o JSON
4. Todas as rotas serão importadas automaticamente!

## 🚨 Troubleshooting

### Swagger não carrega?

1. Verifique se o servidor está rodando: `http://localhost:8080/api/v1/health`
2. Regenere a documentação: `make swagger`
3. Reinicie o servidor: `Ctrl+C` e `make run`

### Mudanças não aparecem?

1. Execute: `make swagger`
2. Recarregue a página do Swagger (F5)
3. Limpe o cache do navegador se necessário (Ctrl+Shift+R)

### Erro 404 no Swagger?

Verifique se o import está correto no main.go:
```go
_ "github.com/misteregis/go-example-project/docs"
```

## ✅ Checklist de Testes

- [ ] Health check retorna status OK
- [ ] GET /users retorna array de usuários
- [ ] GET /users/1 retorna usuário específico
- [ ] POST /users cria novo usuário
- [ ] Validação de nome funciona
- [ ] Validação de email funciona
- [ ] Busca de usuário inexistente retorna 404
- [ ] Swagger UI carrega corretamente
- [ ] Schemas aparecem na documentação
- [ ] cURL gerado funciona no terminal
