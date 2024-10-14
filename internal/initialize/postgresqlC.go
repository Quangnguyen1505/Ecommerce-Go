package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/model"
	"go.uber.org/zap"
	"gorm.io/gen"
)

func CheckErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitPostgresqlC() {
	ctx := context.Background()
	m := global.Config.Postgresql
	fmt.Println("m.Host, m.Username, m.Password, m.Dbname, m.Port::", m.Host, m.Username, m.Password, m.Dbname, m.Port)
	dsn := "user=%s password=%s dbname=%s"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Dbname)
	db, err := pgx.Connect(ctx, s)
	CheckErrorPanic(err, "InitPostgresqlC initialization error")
	global.Logger.Info("initialization PostgresqlC Successfully!")

	global.Pdbc = db
	// defer db.Close(ctx)
	//set pool
	// SetPoolC()
	// genDbs()
	// migrateTablesC()
}

func SetPoolC() {
	m := global.Config.Postgresql
	sqlDb, err := global.Pdb.DB()
	if err != nil {
		fmt.Println("postgresql error::", err)
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))

}

func genDbsC() {
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

func migrateTablesC() {
	err := global.Pdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUserV2{},
	)
	if err != nil {
		fmt.Println("postgresql error::", err)
	}
}
