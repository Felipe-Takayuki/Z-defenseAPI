package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteHosp(ctx *gin.Context) {
	id := ctx.Query("id") 
	zombie := schemas.Zombie{}
	hosp := schemas.Hospedeiro{}
	err := db.First(&zombie, id).Error; if err != nil {
		ctx.JSON(400, gin.H{
			"erro" : "erro em encontrar esta analise",
		})
		return 
	}
	err = db.First(&hosp, id).Error; if err != nil {
		ctx.JSON(400, gin.H{
			"erro" : "erro em encontrar este hospedeiro",
		})
		return
	}
    
	err = db.Delete(&zombie).Error; if err != nil {
		ctx.JSON(400, gin.H{
			"erro" : "erro ao tentar deletar esta analise",
		})
		return
	}
	err = db.Delete(&hosp).Error; if err != nil {
		ctx.JSON(400, gin.H{
			"erro" : "erro ao tentar deletar este hospedeiro", 
		})
		return 
	}
	ctx.JSON(200, gin.H{
		"mensagem" : "Hospedeiro e Analise deletadas",
		"hospedeiro-deletado" : hosp,
		"analise-deletada" : zombie,
	})
}
