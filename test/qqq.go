package test
//
//import "time"
//
//type Topics struct {
//	Id         int        `gorm:"primary_key"`
//	Title      string     `gorm:"not null"`
//	CategoryId int        `gorm:"not null"`
//	Category   Categories `gorm:"foreignkey:CategoryId"` //文章所属分类外键
//
//}
//// 分类
//type Categories struct {
//	Id   int    `gorm:"primary_key"`
//	Name string `gorm:"not null"`
//}
//
//
//type Sort struct {
//	// 分类id
//	SortID string `gorm:"column:sort_id;primary_key"`
//	// 分类名称
//	SortName string `gorm:"type:varchar(20)"`
//	// 用户
//	UserID string `gorm:"type:varchar(32)"`
//	// 用户
//	User user.User `gorm:"foreignkey:UserID"`
//	// 创建时间
//	CreatedAt time.Time
//	// 更新时间
//	UpdatedAt time.Time
//}
//
//// 用户
//type User struct {
//	// 用户id
//	UserID string `gorm:"primary_key"`
//	// 用户名称(账号)
//	UserName string `gorm:"type:varchar(20);not null;unique"`
//	// 密码
//	Password string `gorm:"type:varchar(128);not null"`
//	// 昵称
//	NickName string `gorm:"type:varchar(20);not null;unique"`
//	// 邮箱
//	Email string `gorm:"type:varchar(50);not null"`
//	// 备注
//	Remark string `gorm:"type:varchar(100)"`
//	// 是否是活跃用户
//	IsActive bool `gorm:"not null;default:true"`
//	// 是否是超级用户
//	IsSuperuser bool `gorm:"not null;default:false"`
//	// 创建时间
//	CreatedAt time.Time
//	// 更新时间
//	UpdatedAt time.Time
//}
