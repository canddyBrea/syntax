package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                                                                                                            // 允许访问所有域
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                                                    // 是否允许发送 Cookie
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")                                                                                                                             // 缓存时间
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With") // 设置允许的访问方式
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")                                                                                             // 设置允许的请求方式
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")                                    // 设置允许携带跨域头信息
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
