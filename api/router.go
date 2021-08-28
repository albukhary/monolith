package api

import (
	"github.com/albukhary/monolith/api/handlers"
	"github.com/gin-gonic/gin"
)

func New(h *handlers.Handler) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.GET("/", h.Hello)

	// create user
	r.POST("/user/", h.CreateUser)
	r.GET("/users/", h.GetUsers)
	r.GET("/user/:email/", h.GetUserByEmail)

	return r
}