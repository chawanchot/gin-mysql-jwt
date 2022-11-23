package controller

import (
	"database/sql"
	"fmt"
	"gin-mysql-jwt/dto"
	"gin-mysql-jwt/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserService dto.UserService
	JwtService  dto.JwtService
}

func NewUserController(r *gin.RouterGroup, userService dto.UserService, jwtService dto.JwtService) {
	controller := &userController{
		UserService: userService,
		JwtService:  jwtService,
	}

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/profile", middleware.JwtAuthentication(), controller.Profile)
}

func (u *userController) Register(ctx *gin.Context) {
	req := &dto.UserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.UserResponse{
			Status:  "error",
			Message: "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	user, err := u.UserService.FindOne(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			if err := u.UserService.Create(req); err != nil {
				ctx.JSON(http.StatusBadRequest, &dto.UserResponse{
					Status:  "error",
					Message: err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusCreated, &dto.UserResponse{
				Status:  "success",
				Message: "สมัครสมาชิกสำเร็จ",
			})
			return
		} else {
			fmt.Println(err.Error())
			return
		}
	} else {
		ctx.JSON(http.StatusCreated, &dto.UserResponse{
			Status:  "error",
			Message: "มีชื่อผู้ใช้ " + user.Username + " ในระบบแล้ว",
		})
		return
	}
}

func (u *userController) Login(ctx *gin.Context) {
	req := dto.UserLoginReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.UserResponse{
			Status:  "error",
			Message: "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	user, err := u.UserService.FindOne(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, &dto.UserResponse{
				Status:  "error",
				Message: "ไม่พบผู้ใช้นี้ในระบบ",
			})
			return
		} else {
			fmt.Println(err.Error())
			return
		}
	}

	if req.Password != user.Password {
		ctx.JSON(http.StatusBadRequest, &dto.UserResponse{
			Status:  "error",
			Message: "รหัสผ่านไม่ถูกต้อง",
		})
	}

	token, err := u.JwtService.SignUsersAccessToken(&dto.UsersPassport{
		Id:       user.Id,
		Username: user.Username,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &dto.UsersLoginRes{
		AccessToken: *token,
	})
}

func (u *userController) Profile(ctx *gin.Context) {
	cliams := ctx.MustGet("cliams").(string)

	user, err := u.UserService.FindOne(cliams)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, &dto.UserResponse{
				Status:  "error",
				Message: "ไม่พบผู้ใช้นี้ในระบบ",
			})
			return
		} else {
			fmt.Println(err.Error())
			return
		}
	}

	res := dto.ProfileRes{
		Username:  user.Username,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, &res)
}
