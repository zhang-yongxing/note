package study_note

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"zyx/note/db"
	"zyx/note/utils"
)

func AddNote(c *gin.Context){
	SortID := c.Param("sort_id")
	var AddNoteForm AddNoteForm
	if err := c.ShouldBindJSON(&AddNoteForm); err!=nil{
		c.JSON(422, gin.H{
			"error": err})
		return
	}
	if fErr := AddNoteForm.AddNoteFormVerify(); fErr!=nil{
		c.JSON(422, gin.H{
			"error": fErr})
		return
	}
	var note Note
	note.NoteId = utils.UStr32()
	note.NoteName = AddNoteForm.NoteName
	note.NoteDes = AddNoteForm.NoteDes
	note.NoteContent = AddNoteForm.NoteContent
	note.SortId = SortID
	dbc := db.DB
	dbErr := dbc.Create(&note).Error
	if dbErr!=nil{
		c.Status(500)
		return
	}
	c.JSON(201,gin.H{
		"note_id": note.NoteId,
		"created_time": utils.DatetimeToTimestamp(note.CreatedAt),
	})
}


func AlterNote(c *gin.Context){
	NoteId := c.Param("note_id")
	SortId := c.Param("sort_id")
	var AlterNoteForm AlterNoteForm
	if err := c.ShouldBindJSON(&AlterNoteForm); err!=nil{
		c.JSON(422, gin.H{
			"error": err})
		return
	}
	if fErr := AlterNoteForm.AlterNoteFormVerify(); fErr!=nil{
		c.JSON(422, gin.H{
			"error": fErr})
		return
	}
	dbc := db.DB
	var note Note
	var dbErr error
	dbErr = dbc.Where("note_id = ?", NoteId).Find(&note).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	dbErr = dbc.Model(&note).Update(map[string]interface{}{
		"note_name":AlterNoteForm.NoteName,
		"note_des":AlterNoteForm.NoteDes,
		"note_content":AlterNoteForm.NoteContent,
		"sort_id": SortId,
	}).Error
	if dbErr != nil{
		c.Status(500)
		return
	}
	c.Status(204)
}

// 用户获取的笔记详情
func GetNoteDetail(c *gin.Context)  {
	dbc := db.DB
	var note Note
	var dbErr error
	NoteId := c.Param("note_id")
	dbErr = dbc.Where("note_id = ?", NoteId).Find(&note).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	c.JSON(200, note.NoteToWeb())
}

// 用户获取公开的笔记
func GetPublicNoteDetail(c *gin.Context)  {
	dbc := db.DB
	var note Note
	var dbErr error
	NoteId := c.Param("note_id")
	dbErr = dbc.Where("note_id = ? AND is_deleted = ? And is_hide = ?", NoteId, false, false).Preload("Sort").Find(&note).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	c.JSON(200, note.NoteToWeb())
}

// 用户获取的笔记列表
func GetNoteList(c *gin.Context)  {
	userID := c.Request.Header["user_id"][0]
	dbc := db.DB
	var notes []Note
	dbErr:=dbc.Table("note").
		Select("note.note_id, note.note_name,note.note_content,note.sort_id").
		Joins("left join sort on sort.sort_id = note.sort_id").
		Where("sort.user_id = ?",userID).
		Order("note.create_time desc, note.sort_id").
		Find(&notes).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	nList := make([]map[string]interface{}, len(notes))
	for i,v :=range notes{
		nList[i] = v.NoteToWeb()
	}
	c.JSON(200,nList)
}

// 获取默认笔记列表（首页
func GetNoteDefault(c *gin.Context)  {
	UserID := c.Query("user_id")
	PageNum := c.Query("page_num")
	PageSize := c.Query("page_size")
	SortID:= c.Query("sort_id")
	if UserID == ""{
		c.JSON(422,gin.H{
			"field":"user_id",
			"error":"missing",
		})
		return
	}else if PageNum == ""{
		c.JSON(422,gin.H{
			"field":"page_num",
			"error":"missing",
		})
		return
	}else if PageSize == ""{
		c.JSON(422,gin.H{
			"field":"page_size",
			"error":"missing",
		})
		return
	}
	PageNumInt,errN := strconv.Atoi(PageNum)
	if errN!=nil{

	}
	PageSizeInt,errP := strconv.Atoi(PageSize)
	if errP!=nil{

	}
	OffsetNumInt := (int(PageNumInt)-1) * int(PageSizeInt)
	OffsetNum := strconv.Itoa(OffsetNumInt)
	fmt.Println(PageSize)
	fmt.Println(OffsetNum)
	dbc := db.DB
	var notes []NoteWeb
	where := ""
	if SortID == ""{
		where = "sort.user_id = '"+UserID +"'"
	}else {
		where = "sort.user_id = '"+UserID+"' AND sort.sort_id = '"+SortID+"'"
	}
	fmt.Println(where)
	//var notesC []NoteWeb
	dbErr:=dbc.Table("note").
		Select("note.note_id, note.note_name,note.note_content,note.sort_id,note.note_des,sort.sort_name, note.create_time").
		Joins("left join sort on sort.sort_id = note.sort_id").
		Where(where).
		Order("note.create_time desc").
		Limit(PageSize).
		Offset(OffsetNum).
		Find(&notes).
		Error
	var count int
	dbc.Table("note").
		Select("note.note_id, note.note_name,note.note_content,note.sort_id,note.note_des,sort.sort_name, note.create_time").
		Joins("left join sort on sort.sort_id = note.sort_id").
		Where("sort.user_id = ?",UserID).Count(&count)
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.JSON(200,make([]map[string]interface{},0))
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	nList := make([]map[string]interface{}, len(notes))
	for i,v :=range notes{
		nList[i] = v.NoteToWeb()
	}
	c.JSON(200,gin.H{
		"note_count":count,
		"notes":nList,
	})
}