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

// GetUsers retorna todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

// GetUser retorna um usuário específico
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

// CreateUser cria um novo usuário
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
