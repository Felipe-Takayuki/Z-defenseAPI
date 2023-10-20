package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func ListHospsHandler(ctx *gin.Context) {

	Hosps := []schemas.Hospedeiro{}
	ctx.Header("Content-type", "application/json")
	err := db.Find(&Hosps).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": "erro db.Find",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"hospedeiros": Hosps,
	})
}
