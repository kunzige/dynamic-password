package main

import (
	_ "dynamic-password/user/db"
	"dynamic-password/user/routes"
)

func main() {

	//engine := db.GetDB()
	r := routes.Init()

	r.Run()
}
