package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type DBCOnfig struct {
	Server   string
	Username string
	Password string
	Database string
}

func NewDBConnection(c *DBCOnfig) (e error) {
	db, e = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", c.Username, c.Password, c.Server, c.Database))
	db.LogMode(true)

	if e != nil {
		fmt.Println(fmt.Sprintf("Failed to Postgres Server: %s@%s", c.Server, c.Database))

		return
	}

	fmt.Println(fmt.Sprintf("Connected to Postgres Server: %s@%s", c.Server, c.Database))

	return
}

func GetDB() *gorm.DB {
	return db
}
