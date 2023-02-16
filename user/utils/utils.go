package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

// 校验邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 随机生成四位验证码
func Createcode() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)) //这里面前面的04v是和后面的1000相对应的
}

func SendEmailcodeUtil(ema string) string {
	code := Createcode() // 获取随机数

	var emails = []string{ema}
	//e1 := *(*[]byte)(unsafe.Pointer(&code))
	e := email.NewEmail()
	e.From = "早睡早起吃早饭 <2334096040@qq.com>"
	e.To = emails
	e.Subject = "动态密码安全认证系统"
	e.HTML = []byte("欢迎使用本系统，您的验证码为：<h1>" + code + "</h1>")
	//e.Text = e1
	e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2334096040@qq.com", "nnsbeaxzuuboecdf",
		"smtp.qq.com"))
	return GetMd5(code)
}

func GetMd5(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// strings compare
func CompareString(a string, b string) bool {
	if strings.Compare(a, b) == 0 {
		return true
	}
	return false
}
