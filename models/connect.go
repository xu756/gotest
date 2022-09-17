/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/9/12 14:00
*  @To:
 */

package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var SqliteDB *gorm.DB

func InitsqLite() {
	db, err := gorm.Open(sqlite.Open("sql-data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("sqlite数据库连接失败")
	}
	fmt.Println("sqlite数据库连接成功")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	SqliteDB = db
	err = SqliteDB.AutoMigrate(&BridgeDeck{}, &BA{}, &BB{}, &BC{}, &BD{}, &BE{}, &BF{}, &Ips{})
	if err != nil {
		fmt.Println("sqlite用户表创建失败")
		return
	}

}

func GetSqlitedb() *gorm.DB {
	sqlDB, err := SqliteDB.DB()
	if err != nil {
		log.Print("connect db server failed.")
		InitsqLite()
	}
	if err := sqlDB.Ping(); err != nil {
		err := sqlDB.Close()
		if err != nil {
			return nil
		}
		InitsqLite()
	}
	return SqliteDB
}
