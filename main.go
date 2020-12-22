package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

type verifyRequest struct {
	Secret string `json:"secret"`
	OTP    string `json:"otp"`
}

func main() {
	engine := gin.Default()

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/verify", func(c *gin.Context) {
		var request verifyRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		ok := totp.Validate(request.OTP, request.Secret)
		if !ok {
			c.JSON(400, "fail")
			return
		}

		c.JSON(200, "success")
	})

	engine.Run(":8080")
}
