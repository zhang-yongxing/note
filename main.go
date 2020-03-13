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
	 gr := r.Group("/api")
	 // 创建用户 停用
	 gr.POST("/users", user.AddUsers)
	 // 用户登录
	 gr.POST("/login", user.LoginUsers)
	 // 获取用户信息
	 gr.GET("/user", middleware.LoginAuthentication(), user.GetOwnUserInfo)
	 // 用户退出
	 gr.POST("/logout", middleware.LoginAuthentication(), user.LogoutUsers)
	 // 修改用户信息 停用
	 gr.PUT("/users/:user_id", middleware.LoginAuthentication(), user.AlterUsers)
	 //r.GET("/test", middleware.LoginAuthentication(), user.Test)
	 //获取分类列表
	 gr.GET("/sorts", middleware.LoginAuthentication(), sort.GetSortList)
	 // 创建分类
	 gr.POST("/sorts", middleware.LoginAuthentication(), sort.AddSort)
	 // 获取默认分类列表
	 gr.GET("/default/sorts", sort.GetDefaultSortList)
	 //// 修改分类
	 //gr.PUT("/sort_id", middleware.LoginAuthentication(), sort.AlterSort)
	 //// 删除分类
	 //gr.DELETE("/sort_id", middleware.LoginAuthentication(), sort.DelSort)
	 // 获取分类详情
	 gr.GET("/sorts/:sort_id", middleware.LoginAuthentication(), sort.GetSortDetail)
	 //r.GET("/test", sort.Test)
	 // 添加笔记
	 gr.POST("/sorts/:sort_id/notes", middleware.LoginAuthentication(), study_note.AddNote)
	 // 修改笔记
	 gr.PUT("/sorts/:sort_id/notes/:note_id", middleware.LoginAuthentication(), study_note.AlterNote)
	 // 获取默认笔记列表(分页)
	 gr.GET("/default/notes", study_note.GetNoteDefault)
	 // 获取笔记详情
	 gr.GET("/notes/:note_id", study_note.GetNoteDetail)
	 // 获取公开笔记详情
	 //gr.GET("/sorts/:sort_id/notes/:note_id/public", study_note.GetPublicNoteDetail)
	 // 获取公开的笔记列表
	 //gr.GET("/users/:user_id/notes", study_note.GetPublicNoteDetail)
	 // 获取用户笔记列表
	 gr.GET("/notes", middleware.LoginAuthentication(), study_note.GetNoteList)

	 err := r.Run("0.0.0.0:8000")
	 if err!=nil{
	 	panic(err)
	 }
}
