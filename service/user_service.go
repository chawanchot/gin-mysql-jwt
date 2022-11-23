package service

import (
	"errors"
	"gin-mysql-jwt/db"
	"gin-mysql-jwt/dto"
)

type userService struct{}

func NewUserService() dto.UserService {
	return &userService{}
}

func (u *userService) Create(req *dto.UserRequest) error {
	query := "INSERT INTO users (username, password, phone) VALUES (?, ?, ?)"
	result, err := db.Conn.Exec(query, req.Username, req.Password, req.Phone)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("cannot insert row")
	}

	return nil
}

func (u *userService) FindOne(username string) (*dto.UserRead, error) {
	if err := db.Conn.Ping(); err != nil {
		return nil, err
	}

	query := "SELECT id, username, password, phone, created_at FROM users WHERE username = ?"
	row := db.Conn.QueryRow(query, username)

	user := dto.UserRead{}
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Phone, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
