package http_controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kbiits/dealls-take-home-test/adapters/http/middlewares"
	"github.com/kbiits/dealls-take-home-test/config"
	jwt_util "github.com/kbiits/dealls-take-home-test/utils/jwt"
)

type Controllers struct {
	AuthController    AuthController
	ProfileController ProfileController
	SwipeController   SwipeController
	JwtUtil           *jwt_util.JwtUtil
}

func RegisterRoutes(router *gin.Engine, cfg *config.Config, controllers Controllers) {
	v1 := router.Group("/api/v1")

	// auth group
	auth := v1.Group("/auth")
	auth.POST("/sign-up", controllers.AuthController.SignUp)
	auth.POST("/login", controllers.AuthController.Login)

	requireAuthV1 := v1.Group("", middlewares.RequireUserAuth(controllers.JwtUtil))

	// profile group
	profile := requireAuthV1.Group("/profiles")
	profile.GET("/me", controllers.ProfileController.GetLoggedInProfile)
	profile.PUT("/me", controllers.ProfileController.CompleteProfile)

	// swipes group
	swipe := requireAuthV1.Group("/swipes")
	swipe.GET("/next", controllers.SwipeController.GetNextProfileToSwipe)

	internal := v1.Group("/internal")
	internal.Use(middlewares.RequireInternalAuth(cfg.InternalConfig))
	internal.PUT("/clear-blooms", controllers.SwipeController.ClearBloomsByDate)

}
