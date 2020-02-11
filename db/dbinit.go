package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "note"
)

func init()  {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = gorm.Open("postgres", psqlInfo)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("%v---数据库连接池连接成功", DB)
}