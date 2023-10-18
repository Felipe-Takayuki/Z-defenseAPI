package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func ListToZombiesHandler(ctx *gin.Context) {
    
	Zombs := []schemas.Zombie{}
	ctx.Header("Content-type", "application/json")
	err := db.Find(&Zombs).Error
	if err != nil {
		
		ctx.JSON(400, gin.H{
			"error" : err.Error(),
		})
		return 
	}


	
	ctx.JSON(200, gin.H{
		"data": Zombs,
	})
}
