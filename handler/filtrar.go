package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func FiltrarAnalise(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	classificacao := ctx.Query("classificacao")
	if classificacao == "" {
		ctx.JSON(400, gin.H{
			"erro": "não pode ser nulo",
		})
		return
	}

	if classificacao == "1" || classificacao == "Zumbi extremamente perigoso" {
		classificacao = "Zumbi extremamente perigoso"
	}
	if classificacao == "2" || classificacao == "Zumbi perigoso" {
		classificacao = "Zumbi perigoso"
	}
	if classificacao == "3" || classificacao == "Zumbi normal" {
		classificacao = "Zumbi normal"
	}
	if classificacao == "4" || classificacao == "Zumbi fraco" {
		classificacao = "Zumbi fraco"
	}
	zombie := []schemas.Zombie{}
	err := db.Where(&schemas.Zombie{
		Classificao: classificacao,
	}).Find(&zombie).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": "erro na busca por esse filtro",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"filtro":     classificacao,
		"resultados": zombie,
	})

}
