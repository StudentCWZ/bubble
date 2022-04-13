/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: bubble
   @Date: 2022/4/13 08:46
*/

package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("Conect mysql server failed, err: %v\n", err)
		panic(err)
	}
	// 程序退出关闭数据库连接
	defer dao.CloseDb()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	// 启动服务
	err = r.Run(":9090")
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
