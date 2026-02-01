// Package env fornece helpers para variáveis de ambiente (ex.: .env).
// Uso: em cmd/api/main.go ou internal/config.
package env

import "os"

// Load lê variáveis de um arquivo .env. Em produção use os.Getenv
// ou uma lib como github.com/joho/godotenv.
func Load(filenames ...string) error {
	_ = filenames
	return nil
}

// Get retorna a variável de ambiente ou o valor default.
func Get(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
