package initialize

import (
	"fmt"
	"time"

	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
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
	// genDbs()
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

func genDbs() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Pdb) // reuse your gorm db

	g.GenerateModel("go_crm_user")
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}

func migrateTables() {
	err := global.Pdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUserV2{},
	)
	if err != nil {
		fmt.Println("postgresql error::", err)
	}
}
