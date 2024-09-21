package global

import (
	"github.com/jackc/pgx/v5"
	"github.com/ntquang/ecommerce/pkg/logger"
	"github.com/ntquang/ecommerce/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Pdb    *gorm.DB
	Pdbc   *pgx.Conn
	Redis  *redis.Client
)
