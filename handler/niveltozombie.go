package handler

import (
	"net/http"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func AnaliseToZombie(ctx *gin.Context) {

	id := ctx.Query("id")
	hosp := schemas.Hospedeiro{}
	if err := db.First(&hosp, id).Error; err != nil {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(400, gin.H{
			"mensage": err.Error(),
		})
	}

	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"analise-to-zombie": ToZombie(hosp),
	})

}

/// classificações e requisitos
/*
 1 <= idade  <= 17 (+10), imc >=
*/
