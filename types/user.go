package types

type User struct {
	ID              int
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

type RegisterUserPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
