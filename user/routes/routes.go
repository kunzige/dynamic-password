package routes

import (
	"dynamic-password/user/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/send_code", controllers.SendEmailCode)
	}

	return r
}
