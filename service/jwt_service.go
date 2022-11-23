package service

import (
	"gin-mysql-jwt/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type jwtService struct{}

func NewJwtService() dto.JwtService {
	return &jwtService{}
}

func (j *jwtService) SignUsersAccessToken(req *dto.UsersPassport) (*string, error) {
	claims := dto.UsersClaims{
		Id:       req.Id,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ecommerce_access_token",
			Subject:   "users_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &signed, nil
}
