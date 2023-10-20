package handler

import (
	"strings"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func AtualizarHosp(ctx *gin.Context) {
	req := HospedeiroReq{}
	zombieAtulzar := schemas.Zombie{}
	analise := schemas.Zombie{}
	ctx.BindJSON(&req)
	id := ctx.Query("id")
	hosp := schemas.Hospedeiro{}

	err := db.First(&hosp, id).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": "erro ao tentar obter hospedeiro",
		})
		return
	}
	err = db.First(&analise, id).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": "erro ao tentar obter análise",
		})
		return
	}
	if req.Nome != "" {
		hosp.Nome = req.Nome
	}
	if req.Idade >= 0 {
		hosp.Idade = req.Idade
	}
	if req.Sexo != "" {
		hosp.Sexo = req.Sexo
	}

	if req.Peso > 0 {
		hosp.Peso = req.Peso
	}

	if req.Altura > 0 {
		hosp.Altura = req.Altura
	}
	sang := strings.ToLower(req.TipSanguineo)
	
	if  req.TipSanguineo != "" {
		hosp.GtsMusical = req.GtsMusical
	}
	if req.GtsMusical != "" {
		hosp.GtsMusical = req.GtsMusical
	}
	if req.PraticaQualEsporte != "" {
		hosp.PraticaEsporte = req.PraticaQualEsporte
	}
	if req.JogoPrefer != "" {
		hosp.JogoPrefer = req.JogoPrefer
	}
 
	zombieAtulzar = ToZombie(hosp)
	analise.Forca = zombieAtulzar.Forca
	analise.Velocidade = zombieAtulzar.Velocidade
	analise.Resistencia = zombieAtulzar.Resistencia
	analise.Classificao = zombieAtulzar.Classificao
    
	if sang == "a+" || sang == "a-" || sang == "b+" || sang == "B-" || sang == "o+" || sang == "o-" || sang == "ab+" || sang == "ab-" {
		analise.TipSanguineo = req.TipSanguineo
	} else {
		ctx.JSON(400, gin.H{
			"erro": "tipo sanguíneo desconhecido",
		})
		return
	}
	err = db.Save(&hosp).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": "erro ao tentar salvar as alterações",
		})
		return
	}
	err = db.Save(&analise).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": "erro ao tentar salvar as alterações",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"mensagem": "Hospedeiro atualizado",
		"analise": schemas.Analise{
			Forca:       analise.Forca,
			Velocidade:  analise.Velocidade,
			Resistencia: analise.Resistencia,
			Classificao: analise.Classificao,
		},
		"hospedeiro": hosp,
	})
}
