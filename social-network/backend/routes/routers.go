package routes

import (
	"social-network/api"
	"social-network/middleware/mycors"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(mycors.Cors())

	user := r.Group("/user")
	{
		user.POST("/login", api.UserAPI{}.Login)
		user.POST("/register", api.UserAPI{}.Register)
	}
	return r
}
