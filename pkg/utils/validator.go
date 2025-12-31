package utils

import (
	"regexp"
	"strings"
)

// IsValidEmail verifica se um email é válido
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsEmpty verifica se uma string está vazia
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// SanitizeString remove espaços extras
func SanitizeString(s string) string {
	return strings.TrimSpace(s)
}
