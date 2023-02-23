package controllers

import (
	"dynamic-password/user/db"
	"dynamic-password/user/result"
	"dynamic-password/user/utils"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

func GenerateRsaKey(c *gin.Context) {
	client := db.GetRedis()
	// 产生新的公钥私钥
	prvKey, pubKey := utils.GenRsaKey()
	// 保存到redis中
	client.Set("priKey", string(prvKey), 0)
	client.Set("pubKey", string(pubKey), 0)

	result.Success(c, "新的公钥私钥已产生", nil)
}

func TestGetCiphertext(c *gin.Context) {
	client := db.GetRedis()

	pubKey, _ := client.Get("pubKey").Result()
	password := c.Query("password")
	ciphertext := utils.RsaEncrypt([]byte(password), []byte(pubKey))
	result.Success(c, "", hex.EncodeToString(ciphertext))

}
