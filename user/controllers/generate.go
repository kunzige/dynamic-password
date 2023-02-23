package controllers

import (
	"dynamic-password/user/db"
	"dynamic-password/user/form"
	"dynamic-password/user/models"
	"dynamic-password/user/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Generate(r *gin.Context) {
	//	随机生成hash然后放进redis进行缓存
	user := form.GenerateForm{}
	if err := r.BindJSON(&user); err != nil {
		r.Error(fmt.Errorf("参数不全"))
		return
	}

	temporaryPassword := utils.GetMd5(utils.Createcode())
	redis := db.GetRedis()
	sql := db.GetDB()
	var users []models.User
	_ = sql.SQL("select * from user where email= ? and identity = ?", user.Email, user.Identity).Find(&users)
	if len(users) == 0 {
		r.JSON(403, map[string]interface{}{
			"message": "请先注册",
		})
		return
	}

	redis.Set(user.Identity, temporaryPassword, time.Minute*5)
	r.JSON(200, map[string]interface{}{
		"message":  "成功生成",
		"password": temporaryPassword,
	})
}
