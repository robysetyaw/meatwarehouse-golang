package utils

const (
	// query user
	INSERT      = "INSERT INTO %s (id, username, password, is_active, role, created_at, updated_at, created_by,updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	UPDATE_USER      = "UPDATE users SET username = $2, password = $3, is_active = $4, role = $5, updated_at = $6,updated_by = $7 WHERE id = $1"
	DELETE_USER      = "DELETE FROM users WHERE id = $1"
	GET_USER_BY_ID   = "SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by FROM users WHERE id = $1"
	GET_USER_BY_NAME = "SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by FROM users WHERE username = $1"
	GET_ALL_USER     = "SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by FROM users"
	// query customer
	INSERT_CUSTOMER = "INSERT INTO customer(id, fullname, address, company_id, phone_number, created_At, updated_at, created_by, updated_by) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9aa)"
)
