package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBInit() *gorm.DB {
dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/restfullapi?charset=utf8&parseTime=True&loc=Local")
db,err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err  := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetMaxIdleConns(10)

	return db
}