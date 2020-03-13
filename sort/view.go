package sort

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"zyx/note/db"
	"zyx/note/utils"
)

func GetSortList(c *gin.Context){
	var sorts []Sort
	userID := c.Request.Header["user_id"][0]
	dbc:=db.DB
	dErr := dbc.Where("user_id = ?",userID).Order("create_time desc").Find(&sorts).Error
	if dErr !=nil && dErr.Error() == "record not found" {
		sorts := make([]Sort,0)
		c.JSON(200, sorts)
		return
	}else if dErr !=nil {
		c.Status(500)
		return
	}
	sList := make([]map[string]interface{}, len(sorts))
	for i,v :=range sorts{
		sList[i] = v.SortToWeb()
	}
	c.JSON(200, sList)

}

func GetDefaultSortList(c *gin.Context){
	UserID := c.Query("user_id")
	if UserID == "" {
		c.JSON(422, gin.H{
			"field": "user_id",
			"error": "missing",
		})
		return
	}
	dbc := db.DB
	var sorts []Sort
	dbErr:=dbc.Where("user_id = ?", UserID).Find(&sorts).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		sBlankList:=make([]map[string]interface{}, 0)
		c.JSON(200, sBlankList)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	sList := make([]map[string]interface{}, len(sorts))
	for i, s := range sorts{
		var count int
		dbc.Table("note").Select("sort_id").
			Where("sort_id = ?", s.SortID).Count(&count)
		m := s.SortToWeb()
		m["sort_count"] = count
		sList[i] = m
	}
	c.JSON(200, sList)
}

func GetSortDetail(c *gin.Context)  {
	dbc := db.DB
	var sort Sort
	var dbErr error
	SortId := c.Param("sort_id")
	dbErr = dbc.Where("sort_id = ?", SortId).Find(&sort).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	c.JSON(200, sort.SortToWeb())
}

func AddSort(c *gin.Context) {
	userID := c.Request.Header["user_id"][0]
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
	sort.UserID = userID
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