package router

import (
	"gin-ranking/config"
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.Use((gin.LoggerWithConfig(logger.LoggerToFile())))
	r.Use(logger.Recover)

	store, err := redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfoByID)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUserName)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)

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

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/rank1", controllers.PlayerController{}.GetPlayersRank)
	r.POST("/rank2", controllers.PlayerController{}.GetRankingByRedis)
	return r
}
