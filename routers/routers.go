/*
   @Author: StudentCWZ
   @Description:
   @File: routers
   @Software: GoLand
   @Project: bubble
   @Date: 2022/4/13 08:50
*/

package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建一个默认的路由实例
	r := gin.Default()
	// 告诉 gin 框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// GET 请求
	// 告诉 gin 框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
