package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *sql.DB

var GDB *gorm.DB

//账号数据库

func init() {

	gdb, err := gorm.Open("mysql", "root:123456@(1.2.3.4:22809)/TelAlertMid?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm conn is fail...")
		panic(err)
	}
	// gdb.SingularTable(true)
	gdb.AutoMigrate(&CalledShowTelNum{}, &CalledTelNum{}, &AlertMsg{}, &Voicetem{})

	GDB = gdb

}
