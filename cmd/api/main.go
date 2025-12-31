package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/misteregis/go-example-project/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado")
	}

	// Configurar roteador
	r := mux.NewRouter()

	// Rotas da API
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	api.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	api.HandleFunc("/health", healthCheck).Methods("GET")

	// Porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Servidor funcionando corretamente",
	}); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}
