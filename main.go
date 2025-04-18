package main

import (
	"bubble_front/dao"
	"bubble_front/models"
	"bubble_front/router"
)

func main() {
	
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := router.SetupRouter()
	// // 模型绑定
	// dao.DB.AutoMigrate(&models.Todo{})
	r.Run()
	
}
