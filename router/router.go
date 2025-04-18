package router

import (
	"bubble_front/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Index)
	
	v1Group := r.Group("/v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		
		// 修改待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		
		// 我感觉查看某一个在一个具体的功能里面只能看是否又该功能，不然就是一个基础类型
		// // 查看某一个
		// v1Group.GET("/todo/:id", controller.GetATodoList)
		
		// 删除某一个
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
