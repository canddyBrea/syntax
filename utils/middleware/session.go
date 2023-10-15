package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionMiddleware struct{}

func (*SessionMiddleware) CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		way := c.Request.URL.Path
		if way == "/user/login" || way == "/user/signup" {
			return
		}
		session := sessions.Default(c)
		if session.Get("session_id") == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
