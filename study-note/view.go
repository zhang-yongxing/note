package study_note

import (
	"github.com/gin-gonic/gin"
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
	c.Status(204)
}


func AlterNote(c *gin.Context){
	NoteId := c.Param("note_id")
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
	}).Error
	if dbErr != nil{
		c.Status(500)
		return
	}
	c.Status(204)
}

// 用户获取自己的笔记
func GetNoteDetail(c *gin.Context)  {
	dbc := db.DB
	var note Note
	var dbErr error
	NoteId := c.Param("note_id")
	dbErr = dbc.Where("note_id = ? AND is_deleted = ?", NoteId, false).Preload("Sort").Find(&note).Error
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
	//UserId := c.Param("user_id")
	//dbc := db.DB
	//var sorts []sort.Sort
	////var note Note
	//dbErr := dbc.Where("user_id = ?", UserId).Preload("Note").Find(&sorts).Error
	//fmt.Println(sorts)
	//if dbErr != nil && dbErr.Error() == "record not found"{
	//	c.Status(404)
	//	return
	//}else if dbErr != nil{
	//	c.Status(500)
	//	return
	//}
	//sorts := make([]sort.Sort,0)
	//fmt.Println(sorts)
	//c.JSON(200, sorts)
}