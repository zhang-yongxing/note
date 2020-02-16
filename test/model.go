package test

import (
	"fmt"
	"zyx/note/db"
	"zyx/note/test1"
)

//import (
//	"fmt"
//	"zyx/note/db"
//)
//
type Topics struct {
	Id         int        `gorm:"primary_key"`
	Title      string     `gorm:"not null"`
	CategoryId int        `gorm:"not null"`
	Category   Categories `gorm:"foreignkey:CategoryId"` //文章所属分类外键

}


// 分类
type Categories struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

type Houses struct {
	HouseId string `gorm:"primary_key"`
	People test1.Peoples `gorm:"foreignkey:PeopleId;AssociationForeignKey:PeopleId"`
	PeopleId string
}
//
func Test111(){
	var topic Topics
	dbc := db.DB
	err := dbc.Where("title=?", "wwqwqwqe").Preload("Category").
		Find(&topic).Error

	fmt.Println(err)
	fmt.Println(topic)
	fmt.Println(topic.Category)

	//var cate Categories
	//dbc := db.DB
	//err := dbc.Where("id = ?", 11).Preload("Topics").Find(&cate).Error
	//
	//fmt.Println(err)
	//fmt.Println(cate.Topics)
}