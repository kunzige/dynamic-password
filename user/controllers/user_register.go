package controllers

import (
	"dynamic-password/user/db"
	"dynamic-password/user/form"
	"dynamic-password/user/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	engine := db.GetDB()

	var body form.RegisterForm
	//绑定json和结构体
	if err := c.BindJSON(&body); err != nil {
		return
	}

	user := models.User{Email: body.Email, Password: body.Password}
	_, err := engine.Insert(&user)
	if err != nil {
		fmt.Println("插入失败，", err)
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data":    "注册成功",
		"message": "index",
	})

}
