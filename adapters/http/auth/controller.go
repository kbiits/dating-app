package http_auth

import (
	"github.com/gin-gonic/gin"
	http_controllers "github.com/kbiits/dealls-take-home-test/adapters/http"
	auth_usecase "github.com/kbiits/dealls-take-home-test/usecases/auth"
	validator_util "github.com/kbiits/dealls-take-home-test/utils/validator"
)

type authController struct {
	authUsecase auth_usecase.AuthUsecase
}

func NewAuthController(
	authUsecase auth_usecase.AuthUsecase,
) http_controllers.AuthController {
	return &authController{
		authUsecase: authUsecase,
	}
}

func (a *authController) SignUp(c *gin.Context) {
	req := SignUpRequest{}
	if err := c.BindJSON(&req); err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	if err := validator_util.GetValidator().Struct(&req); err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	signUpResult, err := a.authUsecase.SignUp(c, auth_usecase.SignUpSpec{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(201, http_controllers.NewSuccessResponse(signUpResult))
}

func (a *authController) Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	loginResult, err := a.authUsecase.Login(c, auth_usecase.LoginSpec{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http_controllers.HandleError(c, err)
		return
	}

	c.JSON(200, http_controllers.NewSuccessResponse(loginResult))
}
