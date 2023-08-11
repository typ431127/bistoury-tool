package core

import (
	"bistoury-sync/server/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.GL_MYSQL_USER, global.GL_MYSQL_PASS, global.GL_MYSQL_HOST, global.GL_MYSQL_DB)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		global.GL_LOG.Fatalln("mysql connection failed")
	} else {
		global.GL_LOG.Info("mysql connection is successful")
	}
	return db
}
