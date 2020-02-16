package sort

import (
	"time"
	"zyx/note/user"
	"zyx/note/utils"
)

type Sort struct {
	// 分类id
	SortID string `gorm:"type:varchar(32);primary_key"`
	// 分类名称
	SortName string `gorm:"type:varchar(20)"`
	// 用户
	UserID string `gorm:"type:varchar(32);not null"`
	// 用户
	User user.User `gorm:"foreignkey:UserID;AssociationForeignKey:UserID"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null;column:create_time"`
	// 更新时间
	UpdatedAt time.Time `gorm:"not null;column:update_time"`
}


func (Sort) TableName() string {
	return "sort"
}


func (s Sort)SortToWeb() map[string]interface{}{
	m := make(map[string]interface{}, 4)
	m["sort_id"] = s.SortID
	m["sort_name"] = s.SortName
	m["create_time"] = utils.DatetimeToTimestamp(s.CreatedAt)
	m["update_time"] = utils.DatetimeToTimestamp(s.UpdatedAt)
	return m
}