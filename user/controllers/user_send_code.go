package controllers

import (
	"dynamic-password/user/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SendEmailCode(c *gin.Context) {
	type Body struct {
		Email string `json:"email"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		println(err)
		return
	}
	fmt.Println(body.Email)
	if !utils.VerifyEmailFormat(body.Email) {
		c.JSON(400, gin.H{
			"data": nil,
			"msg":  "请输入正确的邮箱",
		})
		return
	}

	_ = utils.SendEmailcodeUtil(body.Email)

	c.JSON(200, gin.H{
		"data": nil,
		"msg":  "验证码发送成功，5分钟内有效",
	})
}
