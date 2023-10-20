package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/config"
	"gorm.io/gorm"
)

var (
    
	db *gorm.DB
)


 func IniciarHandler(){
	
	db = config.GetSQLite()
 }