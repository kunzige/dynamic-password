package controllers

import (
	"dynamic-password/user/db"
	"dynamic-password/user/models"
	"dynamic-password/user/result"
	"dynamic-password/user/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	// 校验邮箱格式
	if !utils.VerifyEmailFormat(body.Email) {
		c.JSON(400, gin.H{
			"data": nil,
			"msg":  "请输入正确的邮箱",
		})
		return
	}

	engine := db.GetDB()

	var users []models.User
	_ = engine.SQL("select * from user where email = ?", body.Email).Find(&users)
	if len(users) > 0 {
		result.Error(c, http.StatusBadRequest, "该邮箱已存在，不可重复注册")
		return
	}

	// 发送验证码
	code := utils.SendEmailcodeUtil(body.Email)

	// 将验证码保存到redis中
	client := db.GetRedis()
	key := body.Email + "_code"
	err := client.Set(key, code, 5*time.Minute).Err()
	if err != nil {
		fmt.Println(err)
	}

	result.Success(c, "验证码发送成功，5分钟内有效", nil)

}
