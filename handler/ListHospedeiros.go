package handler

import (
	"net/http"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func ListHospsHandler(ctx *gin.Context) {

	Hosps := []schemas.Hospedeiro{}
	db.Find(&Hosps)

	
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": Hosps,
	})
}
