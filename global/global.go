package global

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ntquang/ecommerce/pkg/logger"
	"github.com/ntquang/ecommerce/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Pdb    *gorm.DB
	Pdbc   *pgxpool.Pool
	Redis  *redis.Client
)
