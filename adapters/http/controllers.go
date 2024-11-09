package http_controllers

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type ProfileController interface {
	GetLoggedInProfile(c *gin.Context)
	CompleteProfile(c *gin.Context)
}

type SwipeController interface {
	SwipeProfile(c *gin.Context)
	GetNextProfileToSwipe(c *gin.Context)
	ClearBloomsByDate(c *gin.Context)
}
