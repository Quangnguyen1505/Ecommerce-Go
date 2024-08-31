package initialize

import (
	"fmt"
	"time"

	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CheckErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitPostgresql() {
	m := global.Config.Postgresql
	fmt.Println("m.Host, m.Username, m.Password, m.Dbname, m.Port::", m.Host, m.Username, m.Password, m.Dbname, m.Port)
	dsn := "host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai"
	var s = fmt.Sprintf(dsn, m.Host, m.Username, m.Password, m.Dbname, m.Port)
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	CheckErrorPanic(err, "InitPostgresql initialization error")
	global.Logger.Info("initialization Postgresql Successfully!")

	global.Pdb = db

	//set pool
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Postgresql
	sqlDb, err := global.Pdb.DB()
	if err != nil {
		fmt.Println("postgresql error::", err)
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))

}

func migrateTables() {
	err := global.Pdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("postgresql error::", err)
	}
}
