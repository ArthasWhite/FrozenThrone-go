package global

import (
	"github.com/Arthaslixin/FrozenThrone-go/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
)
