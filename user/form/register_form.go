package form

type RegisterForm struct {
	Email      string `json:"email"`
	Code       string `json:"code"`
	Password   string `json:"password"`
	OkPassword string `json:"okpassword"`
}
