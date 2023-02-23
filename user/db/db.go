package db

import (
	"dynamic-password/user/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
)

var Engine *xorm.Engine

func init() {
	viper.SetConfigName("env")    // 文件名
	viper.SetConfigType("toml")   // 文件类型
	viper.AddConfigPath("config") // 路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println(viper.Get("mysql.username"))

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		viper.Get("mysql.userName"), viper.Get("mysql.password"),
		viper.Get("mysql.ipAddress"), viper.Get("mysql.port"),
		viper.Get("mysql.dbName"), viper.Get("mysql.charset"))
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	err = engine.Sync(new(models.User))
	if err != nil {
		fmt.Println("表创建失败")
	}

	Engine = engine

}

func GetDB() *xorm.Engine {
	return Engine
}
