package dto

type UserService interface {
	Create(req *UserRequest) error
	FindOne(username string) (*UserRead, error)
}

type ProfileRes struct {
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserRead struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}
