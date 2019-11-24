package cli

import (
	"fmt"
	"log"
	"os"
	"seed/conf"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func InitializeDB() {
	var err error
	if conf.Config.DataDir == "" {
		conf.Config.DataDir = "."
	}
	if err := os.MkdirAll(conf.Config.DataDir, 0755); err != nil {
		panic(err)
	}
	dbFile := conf.Config.DataDir + "/seed.db"
	fmt.Println("db_file:", dbFile)
	DB, err = gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Println(err)
		panic("连接数据库失败")
	}
	DB.LogMode(true)
}
