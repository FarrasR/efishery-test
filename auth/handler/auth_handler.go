package handler

import (
	"efishery-auth/entity/form"
	"efishery-auth/entity/response"
	"efishery-auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Register(router *gin.Engine) {
	router.POST("/register", h.RegisterUser)
	router.POST("/login", h.Login)
	router.GET("/explain", h.ExplainClaim)
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var form form.RegisterForm

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	if form.Phone == "" || form.Role == "" || form.Username == "" {
		response.ErrorInvalidParameter(c)
		return
	}

	password, err := h.AuthService.Register(form, c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	registerResponse := response.RegisterResponse{
		Username: form.Username,
		Phone:    form.Phone,
		Role:     form.Role,
		Password: *password,
	}

	response.OKWithHTTPCode(c, http.StatusCreated, "User Created Successfully", registerResponse)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var form form.LoginForm

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ErrorInvalidParameter(c)
		return
	}

	if form.Phone == "" || form.Password == "" {
		response.ErrorInvalidParameter(c)
		return
	}

	token, err := h.AuthService.Login(form, c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	loginResponse := response.LoginResponse{
		Token: *token,
	}

	response.OK(c, loginResponse)
}

func (h *AuthHandler) ExplainClaim(c *gin.Context) {
	token := c.GetHeader("Authorization")

	claim, err := h.AuthService.Explain(token, c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.OK(c, claim)
}
