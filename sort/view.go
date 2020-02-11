package sort

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	sort.UserID = c.Request.Header["user_id"][0]
	sort.UpdatedAt = time.Now()
	sort.CreatedAt = time.Now()
	dbc.Create(&sort)
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
	dbc.Where("sort_name = ? and user_id = ?", ALterSortForm.SortName, UserID).Count(&count)
	dbc.Where("sort_id = ? ", SortID).Find(&sort)
	dbc.Model(&sort).Update(map[string]interface{}{"sort_name": ALterSortForm.SortName})
	c.Status(204)
}

func DelSort(c *gin.Context){
	SortID := c.Param("sort_id")
	dbc := db.DB
	dbc.Where("sort_id = ?", SortID).Delete(Sort{})
	c.Status(204)
}
