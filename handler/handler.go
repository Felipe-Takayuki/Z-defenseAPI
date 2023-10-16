package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/config"
	"gorm.io/gorm"
)

var (
    logger *config.Logger
	db *gorm.DB
)


 func IniciarHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
 }