package routes

import (
	"social-network/api"
	"social-network/config"
	"social-network/middleware/logger"
	"social-network/middleware/mycors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	r.Use(mycors.Cors())

	store, err := redis.NewStore(10, "tcp", config.RedisHost, "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		user.POST("/login", api.UserAPI{}.Login)
		user.POST("/register", api.UserAPI{}.Register)
		user.POST("/modifypassword", api.UserAPI{}.ModifyPWD)
		user.POST("/list", api.UserAPI{}.GetAllUserList)
		user.POST("/pagelist", api.UserAPI{}.GetPageUserList)
		user.POST("/delete", api.UserAPI{}.DeleteUser)
	}
	r.Static("/uploads", "./uploads")
	moment := r.Group("/moment")
	{
		moment.POST("/send", api.MomentAPI{}.SendMoment)
		moment.POST("/uploadfile", api.MomentAPI{}.UploadFile)
		moment.POST("/pagelist", api.MomentAPI{}.GetPageMomentList)
		moment.POST("/usermoments", api.MomentAPI{}.GetMomentsByUserID)
		moment.POST("/ranked", api.MomentAPI{}.RankMoments)
		moment.POST("/delete", api.MomentAPI{}.DeleteMoment)
		moment.POST("/like", api.LikeAPI{}.LikeMomentByRedis)
		moment.POST("/unlike", api.LikeAPI{}.UnLikeMomentByRedis)

	}

	return r
}
