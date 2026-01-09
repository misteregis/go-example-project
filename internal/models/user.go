package models

import "time"

// User representa um usuário do sistema
type User struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"João Silva"`
	Email     string    `json:"email" example:"joao@example.com"`
	Age       int       `json:"age" example:"30"`
	CreatedAt time.Time `json:"created_at" example:"2026-01-09T10:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2026-01-09T10:00:00Z"`
}

// Validate valida os dados do usuário
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrInvalidName
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	return nil
}

// Erros personalizados
var (
	ErrInvalidName  = NewValidationError("nome inválido")
	ErrInvalidEmail = NewValidationError("email inválido")
)

// ValidationError representa um erro de validação
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// NewValidationError cria um novo erro de validação
func NewValidationError(msg string) error {
	return &ValidationError{Message: msg}
}
