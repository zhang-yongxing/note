package main

import (
	"github.com/gin-gonic/gin"
	"zyx/note/middleware"
	"zyx/note/sort"
	study_note "zyx/note/study-note"
	"zyx/note/user"
	"zyx/note/utils"
)

func main()  {
	 //dbc:=db.DB
	 ////dbc.CreateTable(user.User{})
	 ////dbc.CreateTable(sort.Sort{})
	 //dbc.CreateTable(test.Categories{})
	 //dbc.CreateTable(test.Topics{})
	 //dbc.CreateTable(test1.Peoples{})
	 //dbc.CreateTable(test.Houses{})
	 //func(){
		// dbc:=db.DB
		// var p test1.Peoples
		// var h test.Houses
		// p.Name = "xiaoming"
		// p.PeopleId = "1"
		// h.HouseId = "12"
		// h.PeopleId = "1"
		// e1 := dbc.Create(&p).Error
		// e2 := dbc.Create(&h).Error
		// fmt.Println(e1)
		// fmt.Println(e2)
		// var house test.Houses
		// //err := dbc.Where("title=?", "wwqwqwqe").Preload("Category").
		// //	Find(&topic).Error
		// err := dbc.Where("house_id=?", "12").Preload("People").
		//	Find(&house).Error
	 //
		// fmt.Println(err)
		// fmt.Println(house)
		// fmt.Println(house.People)
	 //}()

	 r := gin.Default()
	 r.Use(utils.Cors())
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
	 //r.GET("/test", sort.Test)
	 r.POST("/sorts/:sort_id/notes", middleware.LoginAuthentication(), study_note.AddNote)
	 // 修改笔记
	 r.PUT("/sorts/:sort_id/notes/:note_id", middleware.LoginAuthentication(), study_note.AlterNote)
	 // 获取自己的笔记详情
	 r.GET("/sorts/:sort_id/notes/:note_id", middleware.LoginAuthentication(), study_note.GetNoteDetail)
	 // 获取公开笔记详情
	 r.GET("/sorts/:sort_id/notes/:note_id/public",study_note.GetPublicNoteDetail)
	 // 获取公开的笔记列表
	 r.GET("/sorts/:sort_id/notes",study_note.GetPublicNoteDetail)
	 err := r.Run("0.0.0.0:8000")
	 if err!=nil{
	 	panic(err)
	 }
}
