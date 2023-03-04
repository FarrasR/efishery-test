package form

type RegisterForm struct {
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
