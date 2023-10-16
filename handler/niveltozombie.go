package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func AnaliseToZombie(ctx *gin.Context) {
	id:= ctx.Query("id")
    hosp := schemas.Hospedeiro{}
	if err := db.First(&hosp, id ).Error; err != nil {
		ctx.Header("Content-type", "application/json")
	ctx.JSON(400, gin.H{
		"mensage":   err.Error(),
	})
	}
	if (hosp.Idade >= 18 && hosp.Idade <= 20) && (hosp.PraticaEsporte != "eSport" || hosp.PraticaEsporte != "Nenhum")  {
		ctx.Header("Content-type", "application/json")
	ctx.JSON(200, gin.H{
		"classificação" : "A decidir",
	})
	}
}