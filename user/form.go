package user

// 创建用户
type AddUserForm struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Remark string `json:"remark"`
}

// 修改用户信息
type AlterUserForm struct {
	NickName string `json:"nick_name" binding:"required"`
	Remark string `json:"remark" binding:"required"`
}

type LoginUsersForm struct {
	// 用户名（账号）
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}