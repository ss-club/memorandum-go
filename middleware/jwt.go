package middleware

import (
	"github.com/gin-gonic/gin"
	"gogogo/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 402
			}

		}
		if code != 200 {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "token出错",
				"data":   "请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
