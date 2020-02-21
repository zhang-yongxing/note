package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"zyx/note/utils"
)

func LoginAuthentication() gin.HandlerFunc{

	return func(context *gin.Context) {
		cv, err := context.Request.Cookie("user_id")
		if cv == nil{
			context.JSON(403, "Permission denied")
			context.Abort()
			log.Println("没有cookie", err)
			return
		}
		rv := utils.RedisGet(cv.Value)
		var userSta string
		rErr := rv.Scan(&userSta)
		if rErr != nil{
			context.JSON(403, "Permission denied")
			context.Abort()
			return
		}
		if userSta == ""{
			context.JSON(403, "Permission denied")
			context.Abort()
		}
		sList := make([]string, 1)
		sList[0] = userSta
		cList := make([]string, 1)
		cList[0] = "cv.Value"
		header := context.Request.Header
		header["user_id"] = sList
		header["cookie_userid_key"] = cList
		context.Next()
	}
}

