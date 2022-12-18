package handler

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/hafizdkn/toko-belanja/apps/domain/handler/views"
	"github.com/hafizdkn/toko-belanja/apps/domain/user"
)

type userHandler struct {
	service user.IService
}

func NewUserHandler(service user.IService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) UserRegister(c *gin.Context) {
	var input user.UserRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	user, err := h.service.Create(&input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(user, "Success create user")
	views.WriteJsonRespnse(c, resp)
}

func (h *userHandler) UserLogin(c *gin.Context) {
	var input user.UserLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	token, err := h.service.Login(&input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(token, "Success login")
	views.WriteJsonRespnse(c, resp)
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(users, "Success get users")
	views.WriteJsonRespnse(c, resp)
}

func (h *userHandler) UserUpdateBalance(c *gin.Context) {
	var input user.UserTopUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	userEmail := GetCurrentUserEmail(c)

	err := h.service.UpdateBalance(userEmail, &input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	msg := fmt.Sprintf("Your balance has been successfully update to Rp %d", input.Balance)
	resp := views.SuccessCreateResponse(nil, msg)
	views.WriteJsonRespnse(c, resp)
}

func GetCurrentUserEmail(c *gin.Context) string {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userEmail := userData["email"].(string)
	return userEmail
}

func GetCurrentUserID(c *gin.Context) int {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	return userID
}
