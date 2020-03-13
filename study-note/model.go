package study_note

import (
	"time"
	"zyx/note/sort"
	"zyx/note/utils"
)

type Note struct {
	NoteId string `gorm:"type:varchar(32);not null;primary_key"`
	// 名称
	NoteName string `gorm:"type:varchar(50);not null"`
	// 描述
	NoteDes string `gorm:"type:varchar(50)"`
	// 内容
	NoteContent string `gorm:"not null"`
	// 是隐藏状态
	IsHide bool `gorm:"default:true"`
	// 是删除状态
	IsDeleted bool `gorm:"default:false"`
	// 分类对象
	Sort sort.Sort `gorm:"foreignkey:SortId;AssociationForeignKey:SortId"`
	// 分类id
	SortId string `gorm:"type:varchar(32);not null"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null;column:create_time"`
	// 更新时间
	UpdatedAt time.Time `gorm:"not null;column:update_time"`
}

type NoteWeb struct {
	Note
	SortName string
}

func (Note) TableName() string {
	return "note"
}


func (n Note)NoteToWeb() map[string]interface{}{
	m := make(map[string]interface{}, 9)
	m["note_id"] = n.NoteId
	m["note_name"] = n.NoteName
	m["note_des"] = n.NoteDes
	m["note_content"] = n.NoteContent
	m["sort_id"] = n.SortId
	m["create_time"] = utils.DatetimeToTimestamp(n.CreatedAt)
	m["update_time"] = utils.DatetimeToTimestamp(n.UpdatedAt)
	return m
}

func (n NoteWeb)NoteToWeb() map[string]interface{}{
	m := make(map[string]interface{}, 9)
	m["note_id"] = n.NoteId
	m["note_name"] = n.NoteName
	m["note_des"] = n.NoteDes
	m["note_content"] = n.NoteContent
	m["sort_id"] = n.SortId
	m["sort_name"] = n.SortName
	m["create_time"] = utils.DatetimeToTimestamp(n.CreatedAt)
	m["update_time"] = utils.DatetimeToTimestamp(n.UpdatedAt)
	return m
}

