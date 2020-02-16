package user

import (
	"time"
	"zyx/note/utils"
)

// 用户
type User struct {
	// 用户id
	UserID string `gorm:"type:varchar(32);primary_key"`
	// 用户名称(账号)
	UserName string `gorm:"type:varchar(20);not null;unique"`
	// 密码
	Password string `gorm:"type:varchar(128);not null"`
	// 昵称
	NickName string `gorm:"type:varchar(20);not null;unique"`
	// 邮箱
	Email string `gorm:"type:varchar(50);not null"`
	// 备注
	Remark string `gorm:"type:varchar(100)"`
	// 是否是活跃用户
	IsActive bool `gorm:"not null;default:true"`
	// 是否是超级用户
	IsSuperuser bool `gorm:"not null;default:false"`
	// 创建时间
	CreatedAt time.Time `gorm:"not null;column:create_time"`
	// 更新时间
	UpdatedAt time.Time `gorm:"not null;column:update_time"`
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
