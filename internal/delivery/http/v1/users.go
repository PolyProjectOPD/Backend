package v1

import (
	"github.com/PolyProjectOPD/Backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refresh)
	}
}

type signUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary User SignUp
// @Tags User Auth
// @Description User sign-up
// @ModuleID signUp
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "Sign-up info"
// @Success 201 {object} signUpResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Users.SignUp(service.UserSignUpInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "user already exists")
		return
	}

	c.JSON(http.StatusCreated, signUpResponse{
		ID: id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary User SignIn
// @Tags User Auth
// @Description User sign-in
// @ModuleID signIn
// @Accept  json
// @Produce  json
// @Param input body signInInput true "Sign-in info"
// @Success 200 {object} tokensResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.services.Users.SignIn(service.UserSignInInput{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid email or password")
		return
	}

	c.JSON(http.StatusOK, tokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

// @Summary User Refresh Tokens
// @Tags User Auth
// @Description User refresh tokens
// @Accept  json
// @Produce  json
// @Param input body refreshInput true "Refresh tokens info"
// @Success 200 {object} tokensResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.services.Users.RefreshTokens(input.Token)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "invalid refresh token")
		return
	}

	c.JSON(http.StatusOK, tokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
