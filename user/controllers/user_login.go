package controllers

import (
	"dynamic-password/user/db"
	"dynamic-password/user/form"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(r *gin.Context) {
	//	 根据返回的密码进行登录
	loginForm := form.LoginForm{}
	client := db.GetRedis()

	if err := r.BindJSON(&loginForm); err != nil {
		fmt.Printf("%v\n", err)
		r.JSON(400, map[string]interface{}{
			"message": "参数不全",
		})
		return
	}

	userPassword, _ := client.Get(loginForm.Identity).Result()
	if userPassword == "" {
		r.JSON(403, map[string]interface{}{
			"message": "无效的",
		})
		return
	}

	if loginForm.Password != userPassword {
		r.JSON(403, map[string]interface{}{
			"message": "密码错误",
		})
		return
	}
	r.JSON(200, map[string]interface{}{
		"message": "登陆成功",
		"token":   "llllllllllll",
	})
}
