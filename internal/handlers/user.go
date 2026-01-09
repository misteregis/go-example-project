package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/misteregis/go-example-project/internal/models"

	"github.com/gorilla/mux"
)

// Simulação de banco de dados em memória
var users = []models.User{
	{
		ID:        1,
		Name:      "João Silva",
		Email:     "joao@example.com",
		Age:       30,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Maria Santos",
		Email:     "maria@example.com",
		Age:       25,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// GetUsers godoc
// @Summary Listar todos os usuários
// @Description Retorna a lista de todos os usuários cadastrados
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {string} string "Erro ao codificar resposta"
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

// GetUser godoc
// @Summary Buscar usuário por ID
// @Description Retorna um usuário específico pelo ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID do Usuário"
// @Success 200 {object} models.User
// @Failure 400 {string} string "ID inválido"
// @Failure 404 {string} string "Usuário não encontrado"
// @Failure 500 {string} string "Erro ao codificar resposta"
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

// CreateUser godoc
// @Summary Criar novo usuário
// @Description Cria um novo usuário no sistema
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Dados do Usuário"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Dados inválidos"
// @Failure 500 {string} string "Erro ao codificar resposta"
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}
