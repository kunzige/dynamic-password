package form

type RegisterForm struct {
	Email      string `json:"email"`
	Code       string `json:"code"`
	Password   string `json:"password"`
	OkPassword string `json:"okpassword"`
}

type GenerateForm struct {
	Email    string `json:"email"`
	Identity string `json:"identity"`
}

type LoginForm struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
