package controllers

import (
	"dynamic-password/user/db"
	"dynamic-password/user/form"
	"dynamic-password/user/models"
	"dynamic-password/user/result"
	"dynamic-password/user/utils"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Register(c *gin.Context) {
	engine := db.GetDB()
	client := db.GetRedis()

	var body form.RegisterForm
	//绑定json和结构体
	if err := c.BindJSON(&body); err != nil {
		return
	}

	// 获取私钥解密
	priKey, _ := client.Get("priKey").Result()

	ciphertext, _ := hex.DecodeString(body.Password)
	password, flag := utils.RsaDecrypt(ciphertext, []byte(priKey))
	if !flag {
		fmt.Println(password)
		result.Error(c, 400, "密码错误,请联系管理员")
		return
	}
	ciphertext, _ = hex.DecodeString(body.OkPassword)
	okpassword, flag1 := utils.RsaDecrypt(ciphertext, []byte(priKey))
	if !flag1 {
		fmt.Println(okpassword)
		result.Error(c, 400, "确认密码有误,请联系管理员")
		return
	}
	fmt.Println(string(password))
	// 验证两次密码是否一样
	if !utils.CompareString(string(password), string(okpassword)) {
		result.Error(c, http.StatusBadRequest, "两次密码不一致，请重新输入")
		return
	}

	code := body.Code
	key := body.Email + "_code"
	okCode, _ := client.Get(key).Result()

	// 校验验证码
	if !utils.CompareString(utils.GetMd5(code), okCode) {
		result.Error(c, http.StatusBadRequest, "无效验证码")
		return
	}

	// 判断该用户是否注册过
	var users []models.User

	_ = engine.SQL("select * from user where email = ?", body.Email).Find(&users)
	if len(users) > 0 {
		result.Error(c, http.StatusBadRequest, "该邮箱已存在，不可重复注册")
		return
	}

	user := models.User{Email: body.Email, Password: utils.GetMd5(string(password)), Identity: strings.Split(body.Email, "@")[0] + "@" + utils.Createcode()}

	_, err := engine.Insert(&user)
	if err != nil {
		fmt.Println("插入失败，", err)
		result.Error(c, http.StatusInternalServerError, "服务器连接超时")
		return
	}
	result.Success(c, "注册成功", nil)

}
