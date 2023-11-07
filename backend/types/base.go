package types

type ServiceError struct {
	Message string
	Error   error
	Code    int
}
type User struct {
	Id      int  `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type AuthUser struct {
	User
}
