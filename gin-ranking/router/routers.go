package router

import (
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.Use((gin.LoggerWithConfig(logger.LoggerToFile())))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.DELETE("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "用户删除")
		})

	}

	order := r.Group("/order")
	{
		order.GET("/info", controllers.OrderController{}.GetInfo)
		order.POST("/list", controllers.OrderController{}.GetList)
		order.PUT("/add", func(c *gin.Context) {
			c.String(http.StatusOK, "订单添加")
		})
		order.DELETE("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "订单删除")
		})
	}

	return r
}
