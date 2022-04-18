package schema

type ResetPassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type CreateReset struct {
	Email string `json:"email"`
}