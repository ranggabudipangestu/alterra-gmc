package dto

import "github.com/golang-jwt/jwt/v4"

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

type JWTClaims struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
