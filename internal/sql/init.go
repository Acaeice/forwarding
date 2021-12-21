package sql

import (
	"fmt"
	"log"

	"code.meikeland.com/me/forwarding/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	//Db gorm的数据库连接
	Db *gorm.DB
)

// Init 初始化函数
func Init() {
	database := conf.GetDatabase()

	var err error
	Db, err = gorm.Open(mysql.Open(database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect database, datebase config is %v, err is %+v", database, err))
	}
	if database.Debug {
		Db = Db.Debug()
	}

	// Db.AutoMigrate(
	// 	&pkg.SomeStruct{},
	// )
	log.Println("All table AutoMigrate finish.")
}
