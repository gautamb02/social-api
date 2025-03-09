package users

const (
	CreateUserQuery = `INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`
)
