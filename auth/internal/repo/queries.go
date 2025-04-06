package repo

var (
	insertNewUser        = `INSERT INTO users(id, privatekey, username, password, email, phone, role) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	selectUserByUsername = `SELECT id, privatekey, username, password, email, phone, role FROM users WHERE username = $1`
	selectUserByEmail    = `SELECT id, privatekey, username, password, email, phone, role FROM users WHERE email = $1`
)
