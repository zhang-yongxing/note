package sort

import (
	"time"
)

type Sort struct {
	// 分类id
	SortID string `gorm:"column:sort_id"`
	// 分类名称
	SortName string `gorm:"column:sort_name"`
	// 用户
	UserID string `gorm:"column:user_id"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:create_time"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:update_time"`
}


func (Sort) TableName() string {
	return "sort"
}
