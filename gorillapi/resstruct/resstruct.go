package resstruct

// ErrorResponse struct for error message
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// User struct for user
type User struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}

// UserResponse struct for user response
type UserResponse struct {
	Status string `json:"status"`
	Data   []User `json:"data"`
}
