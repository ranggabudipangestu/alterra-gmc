package dto

type RequestLoginDto struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ResponseLoginDto struct {
	Token string `json:"token"`
}

type RequestRegisterDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"Name"`
	Password string `json:"password"`
}

type ResponseUserDto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
