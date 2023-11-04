package forms

type RegisterPayload struct {
	Username string `json:"username" form:"username"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginPayload struct {
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}