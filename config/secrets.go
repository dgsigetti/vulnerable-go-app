package config

const (
	APIKey          = "hardcoded-secret-key" // Vulnerability: Hardcoded secret
	DatabaseConnStr = "user:password@tcp(localhost:3306)/dbname" // Vulnerability: Plaintext credentials
)

