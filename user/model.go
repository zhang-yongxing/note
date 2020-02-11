package user

import (
	"time"
	"zyx/note/utils"
)

// 用户
type User struct {
	// 用户id
	UserID string `gorm:"column:user_id"`
	// 用户名称(账号)
	UserName string `gorm:"column:user_name"`
	// 密码
	Password string `gorm:"column:password"`
	// 昵称
	NickName string `gorm:"column:nick_name"`
	// 邮箱
	Email string `gorm:"column:email"`
	// 备注
	Remark string `gorm:"column:remark"`
	// 是否删除
	IsActive bool `gorm:"column:is_active"`
	// 是否是超级用户
	IsSuperuser bool `gorm:"column:is_superuser"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:create_time"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:update_time"`
}

func (User) TableName() string {
	return "user"
}

func (u User)UserToWeb() map[string]interface{}{
	m := make(map[string]interface{}, 9)
	m["user_id"] = u.UserID
	m["user_name"] = u.UserName
	m["nick_name"] = u.NickName
	m["email"] = u.Email
	m["remark"] = u.Remark
	m["is_active"] = u.IsActive
	m["is_superuser"] = u.IsSuperuser
	m["create_time"] = utils.DatetimeToTimestamp(u.CreatedAt)
	m["update_time"] = utils.DatetimeToTimestamp(u.UpdatedAt)
	return m
}
