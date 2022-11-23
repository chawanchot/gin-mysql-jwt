package dto

import "github.com/golang-jwt/jwt/v4"

type JwtService interface {
	SignUsersAccessToken(req *UsersPassport) (*string, error)
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersPassport struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

type UsersClaims struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
