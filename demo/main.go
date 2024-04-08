// // middleware/cors_middleware.go
// package middleware

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/rs/cors"
// )

// func CustomCorsMiddleware(config *cors.Cors) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 使用 CORS 配置处理请求
// 		config.Handler(c.Writer, c.Request)

// 		// 如果请求是一个 OPTIONS 请求（预检请求），则直接返回，不继续执行后续的路由处理函数
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 			return
// 		}

//			// 否则，继续执行后续的路由处理函数
//			c.Next()
//		}
//	}
package middleware
