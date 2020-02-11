package main

import (
	"github.com/gin-gonic/gin"
	"zyx/note/middleware"
	"zyx/note/user"
	"zyx/note/sort"
)


func main()  {
	 r := gin.Default()
	 // 创建用户
	 r.POST("/users", user.AddUsers)
	 // 用户登录
	 r.POST("/login", user.LoginUsers)
	 // 获取用户信息
	 r.GET("/user", middleware.LoginAuthentication(), user.GetOwnUserInfo)
	 // 用户退出
	 r.POST("/logout", middleware.LoginAuthentication(), user.LogoutUsers)
	 // 修改用户信息
	 r.POST("/users/:user_id", middleware.LoginAuthentication(), user.AlterUsers)
	 //r.GET("/test", middleware.LoginAuthentication(), user.Test)
	 // 创建分类
	 r.POST("/users/:user_id/sorts", middleware.LoginAuthentication(), sort.AddSort)
	 // 修改分类
	 r.PUT("/users/:user_id/sorts/:sort_id", middleware.LoginAuthentication(), sort.AlterSort)
	 // 删除分类
	 r.DELETE("/users/:user_id/sorts/:sort_id", middleware.LoginAuthentication(), sort.DelSort)
	 err := r.Run("0.0.0.0:8000")
	 if err!=nil{
	 	panic(err)
	 }
}
