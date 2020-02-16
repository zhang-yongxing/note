package sort

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"zyx/note/db"
	"zyx/note/utils"
)

func AddSort(c *gin.Context)  {
	var AddSortForm AddSortForm
	err := c.ShouldBindJSON(&AddSortForm)
	if err != nil{
		errs := fmt.Sprintf("%v", err)
		c.JSON(422, gin.H{
			"error": errs})
		return
	}
	dbc:=db.DB
	var sort Sort
	sort.SortID = utils.UStr32()
	sort.SortName = AddSortForm.SortName
	sort.UserID = c.Param("user_id")
	sort.UpdatedAt = time.Now()
	sort.CreatedAt = time.Now()
	dbErr := dbc.Create(&sort).Error
	if dbErr != nil{
		c.Status(500)
	}
	c.JSON(201, gin.H{
		"sort_id": sort.SortID,
		"created_time": utils.DatetimeToTimestamp(sort.CreatedAt),
	})
}

func AlterSort(c *gin.Context){
	var ALterSortForm ALterSortForm
	if err:= c.ShouldBindJSON(&ALterSortForm); err!=nil{
		errs := fmt.Sprintf("%v", err)
		c.JSON(422, gin.H{
			"error": errs})
		return
	}
	SortID := c.Param("sort_id")
	UserID := c.Param("user_id")
	dbc := db.DB
	var sort Sort
	var count int64
	var dbErr error
	dbErr = dbc.Where("sort_name = ? and user_id = ?", ALterSortForm.SortName, UserID).Count(&count).Error
	if dbErr != nil && dbErr.Error() != "record not found"{
		c.Status(500)
		return
	}
	dbErr = dbc.Where("sort_id = ? ", SortID).Find(&sort).Error
	if dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	dbErr = dbc.Model(&sort).Update(map[string]interface{}{"sort_name": ALterSortForm.SortName}).Error
	if dbErr != nil{
		c.Status(500)
		return
	}
	c.Status(204)
}

func DelSort(c *gin.Context){
	SortID := c.Param("sort_id")
	dbc := db.DB
	dbErr := dbc.Where("sort_id = ?", SortID).Delete(Sort{}).Error
	if dbErr != nil{
		c.Status(500)
		return
	}
	c.Status(204)
}

func Test(c *gin.Context){
	var sort Sort
	//var user1 user.User
	dbc := db.DB
	err := dbc.Where("sort_id = ?", "b1cf829af8ff4113b4371d81efc24c0b").
		Preload("User").Find(&sort).Error


	if err != nil {
		log.Println(err)
	}
	fmt.Println(sort)
	fmt.Println("------------------------------------------------")
	fmt.Println(sort.User)
	c.JSON(201,&sort)
}