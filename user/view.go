package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	"zyx/note/db"
	"zyx/note/utils"
)

// 创建用户
func AddUsers(c *gin.Context){
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//fmt.Println(string(buf[0:n]))

	var userForm AddUserForm
	if err:=c.ShouldBindJSON(&userForm); err!=nil{
		fmt.Println(userForm)
		fmt.Println(err.Error())
		errs := fmt.Sprintf("%v", err)
		fmt.Println(errs)
		c.JSON(422, gin.H{
			"error": errs})
		return
	}
	var user User
	var count int
	var dbErr error
	dbc:=db.DB
	dbErr = dbc.Where("user_name = ?", userForm.UserName).Or("nick_name = ?", userForm.NickName).Find(&user).Count(&count).Error
	if dbErr != nil && dbErr.Error() != "record not found"{
		log.Println(dbErr)
		log.Println(dbErr.Error() == "record not found")
		log.Printf("%T",dbErr.Error())
		c.Status(500)
		return
	}
	if count>= 1{
		if user.UserName == userForm.UserName{
			c.JSON(422, gin.H{
				"field": "user_name",
				"error": "already exists",
			})
			return
		}
		if user.NickName == userForm.NickName{
			c.JSON(422, gin.H{
				"field": "nick_name",
				"error": "already exists",
			})
		}
			return
	}
	pw:=[]byte(userForm.Password)
	hashPassword, err:=bcrypt.GenerateFromPassword(pw,bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(500, gin.H{
			"error": "Password encryption error",
		})
	}
	user.UserID = utils.UStr32()
	user.UserName = userForm.UserName
	user.NickName = userForm.NickName
	user.Password = string(hashPassword)
	user.Email = userForm.Email
	user.Remark = userForm.Remark
	user.IsActive = true
	user.IsSuperuser = false
	user.UpdatedAt = time.Now()
	dbErr = dbc.Create(&user).Error
	if dbErr != nil{
		log.Println(dbErr)
		c.Status(500)
		return
	}
	c.JSON(201, gin.H{
		"user_id": user.UserID,
		"created_time": utils.DatetimeToTimestamp(user.CreatedAt),
	})
}
// 用户登录
func LoginUsers(c *gin.Context)  {
	var loginUsersForm LoginUsersForm
	if err := c.ShouldBindJSON(&loginUsersForm); err != nil{
		c.JSON(422, gin.H{
			"error": err,
		})
		return
	}
	var user User
	var count int
	dbc := db.DB
	dbc.Where( "user_name = ? ", loginUsersForm.UserName).Find(&user).Count(&count)
	if count == 0{	// 用户名不存在
		c.JSON(422, gin.H{
			"error": "user_name or password error",
		})
		return
	}else if count > 1 { // 用户名重复
		c.Status(500)
		return
	}
	// 密码验证
	err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUsersForm.Password))
	if err1!=nil{
		c.JSON(422, gin.H{
			"error": "user_name or password error",
		})
		return
	}
	// 登录成功状态redis存储
	uuidStr := utils.UStr32()
	utils.RedisSet(uuidStr, user.UserID, 60*60*24*30 * time.Second)
	log.Println("设置redis userid:",user.UserID)
	// 设置cookie
	cookie := http.Cookie{
		Name:       "user_id",
		Value:      uuidStr,
		Path:       "",
		Domain:     "www.note.com",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     60*60*24*30,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	http.SetCookie(c.Writer, &cookie)
	//c.SetCookie("user_id", uuidStr, 3600, "/", "localhost", false, true)
	c.JSON(204, "")
}
// 获取登录用户的信息
func GetOwnUserInfo(c *gin.Context)  {
	userID := c.Request.Header["user_id"][0]
	dbc := db.DB
	var user User
	//err := dbc.Where("user_id = ?", userID).Count(count).Find(&user).Error
	err:=dbc.Where("user_id = ?", userID).Find(&user).Error
	if err != nil{
		fmt.Println(err)

	}
	if err != nil{
		c.Status(500)
		fmt.Println(err)
		return
	}

	c.JSON(200, user.UserToWeb())
}
// 用户退出
func LogoutUsers(c *gin.Context){
	cUseridKey := c.Request.Header["cookie_userid_key"][0]
	// 设置cookie为空字符串
	cookie := http.Cookie{
		Name:       "user_id",
		Value:      "",
		Path:       "",
		Domain:     "www.note.com",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     1,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	http.SetCookie(c.Writer, &cookie)
	// redis清空
	intCmd := utils.RedisDel(cUseridKey)
	fmt.Println("intCmd--------:", intCmd)
	c.JSON(204,"")
}
// 修改用户信息
func AlterUsers(c *gin.Context){
	var AlterUserForm AlterUserForm
	userID := c.Param("user_id")
	if err := c.ShouldBindJSON(&AlterUserForm); err != nil{
		c.JSON(422, gin.H{
			"errors": err,
		})
		return
	}
	dbc := db.DB
	var count int64
	var dbErr error
	var cUser User
	dbErr = dbc.Where("nick_name = ? and user_id != ?", AlterUserForm.NickName,userID).Find(&cUser).Count(&count).Error
	if count != 0{
		c.JSON(422, gin.H{
			"errors": "nick_name already exists",
		})
		return
	}
	if dbErr != nil && dbErr.Error() != "record not found"{
		c.Status(500)
		return
	}
	var user User
	dbErr = dbc.Where("user_id = ?", userID).Find(&user).Error
	if dbErr != nil && dbErr.Error() == "record not found"{
		c.Status(404)
		return
	}else if dbErr != nil{
		c.Status(500)
		return
	}
	dbc.Model(&user).Update(map[string]interface{}{"nick_name":AlterUserForm.NickName, "remark":AlterUserForm.Remark})
	c.Status(204)
}

func Test(c *gin.Context){
	//cookie := http.Cookie{
	//	Name:       "user_id",
	//	Value:      "9999",
	//	Path:       "",
	//	Domain:     "",
	//	Expires:    time.Time{},
	//	RawExpires: "",
	//	MaxAge:     15,
	//	Secure:     false,
	//	HttpOnly:   false,
	//	SameSite:   0,
	//	Raw:        "",
	//	Unparsed:   nil,
	//}
	//http.SetCookie(c.Writer, &cookie)
	//v, err := c.Request.Cookie("user_id")
	//
	//fmt.Println(v)
	//fmt.Printf("%T", v.Value)
	//fmt.Printf("类型： %T  ", v)
	//fmt.Println(err)
	//fmt.Printf("%T", err)
	//utils.RedisSet("name","xiaoming",10*time.Second)
	//sCmd := utils.RedisGet("name")
	//fmt.Printf("%T\n", sCmd)
	//fmt.Printf("%T\n", sCmd.String())
	//fmt.Printf("%v\n", sCmd.String())
	//var rv string
	//fmt.Printf("%v\n", sCmd.Scan(&rv))
	//fmt.Printf("%T", rv)
	//c.JSON(201, "ok")
	utils.RedisDel("a")
}