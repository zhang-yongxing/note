package test1

//import (
//	"fmt"
//	"zyx/note/db"
//)
//
//type Topics struct {
//	Id         int        `gorm:"primary_key"`
//	Title      string     `gorm:"not null"`
//	CategoryId int        `gorm:"not null"`
//	Category   Categories `gorm:"foreignkey:CategoryId"` //文章所属分类外键
//}

type Peoples struct {
	PeopleId string `gorm:"primary_key"`
	Name string
}

func Test111(){
	//var p Peoples
	//var h test.Houses
	//p.Name = "xiaoming"
	//p.Id = "1"
	//h.Id = "12"
	//h.PeopleId = "1"
	//dbc := db.DB
	//dbc.Create(p)
	//dbc.Create(h)
	//var house test.Houses
	////err := dbc.Where("title=?", "wwqwqwqe").Preload("Category").
	////	Find(&topic).Error
	//err := dbc.Where("id=?", "1").Preload("People").
	//	Find(&house).Error
	//
	//fmt.Println(err)
	//fmt.Println(house)
	//fmt.Println(house.People)
}
