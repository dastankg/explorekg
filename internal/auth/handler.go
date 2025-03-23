package auth

import (
	"explorekg/config"
	"explorekg/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthHandler struct {
	*config.Config
	*AuthService
}
type AuthHandlerDeps struct {
	*config.Config
	*AuthService
}

func NewAuthHandler(router *gin.Engine, deps AuthHandlerDeps) {
	handler := &AuthHandler{Config: deps.Config, AuthService: deps.AuthService}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
	}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var body LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": ErrWrongCredetials})
		return
	}
	email, err := handler.AuthService.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	jwtService := jwt.NewJWT(
		handler.Config.Auth.AccessToken,
		handler.Config.Auth.RefreshToken,
	)
	tokens, err := jwtService.CreateTokenPair(
		email,
		15*time.Minute,
		24*7*time.Hour,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя и возвращает токены
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Данные для регистрации"
// @Success 201 {object} RegisterResponse "Успешная регистрация"
// @Router /auth/register [post]
func (handler *AuthHandler) Register(c *gin.Context) {
	var body RegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, err := handler.AuthService.Register(body.Name, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	jwtService := jwt.NewJWT(
		handler.Config.Auth.AccessToken,
		handler.Config.Auth.RefreshToken,
	)
	tokens, err := jwtService.CreateTokenPair(
		email,
		15*time.Minute,
		24*7*time.Hour,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
