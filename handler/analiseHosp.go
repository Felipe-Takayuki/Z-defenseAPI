package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func AnaliseHosp(ctx *gin.Context) {
	Hospedeiro := schemas.Hospedeiro{}
	ctx.Header("Content-type", "application/json")
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"mensage" : "o id não pode ser nulo",
		})
		return 
	}

	err := db.First(&Hospedeiro, id).Error; if err != nil {
         ctx.JSON(400, gin.H{
			"mensage" : err.Error(),
		 })
		 return 
	}
	analise := ToZombie(Hospedeiro)
	
	ctx.JSON(200, gin.H{
		"analise" : schemas.Analise{
			Forca: analise.Forca,
			Velocidade: analise.Forca,
			Resistencia: analise.Resistencia,
			Classificao: analise.Classificao,
		},
		"hospedeiro" : Hospedeiro,
	})
}
