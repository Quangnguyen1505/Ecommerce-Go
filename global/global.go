package global

import (
	"github.com/ntquang/ecommerce/pkg/logger"
	"github.com/ntquang/ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Pdb    *gorm.DB
)
