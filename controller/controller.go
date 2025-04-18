package controller

import (
	"bubble_front/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	// 前端页面填写待办事项，加入html内容
	// 1 从请求中把数据拿出
	var todo models.Todo
	c.BindJSON(&todo)
	// 2 存入数据库
	if err := models.CreateATodo(&todo); err != nil { // 把逻辑判断写到逻辑包里面去
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": todo,
		})
	}
	// 3 反应响应
	
}

func GetTodoList(c *gin.Context) {
	todoList, err := models.GetTodoList() // 最好不要给任何参数 ,不明白返回值为什么这样子使用，不知道是不是前段需要这样只返回调用
	if err != nil {                       // ?居然也没有出现报错的情况出现
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// func GetATodoList(c *gin.Context) {
// 	// id := c.Param("id")
// 	id, ok := c.Params.Get("id")
// 	if !ok {
// 		c.JSON(http.StatusOK, gin.H{
// 			"error": "无效的id",
// 		})
// 		return
// 	}
// 	todo, err := models.GetATodo(id)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"error": err.Error(), // record not found
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, todo)
// }

func UpdateATodo(c *gin.Context) {
	// 先用id去查询 param为path参数，query为查询
	id, ok := c.Params.Get("id")
	// id := c.Param("id")
	// fmt.Println("id-------------------" + id)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "参数获取失败",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	// 绑定todo,前面虽然已经给todo赋值了数据库返回的值，但是这里需要给它绑定修改的值
	c.BindJSON(todo)
	if err := models.UpdateAtodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id参数无效",
		})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": "deleted",
	})
}
