package handler

import (
	"strings"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/utils"
	"github.com/gin-gonic/gin"
)

func CatalogHospHandler(ctx *gin.Context) {
	req := HospedeiroReq{}
	ctx.BindJSON(&req)
	ctx.Header("Content-type", "application/json")
	hospedeiro := schemas.Hospedeiro{
		Nome:           req.Nome,
		Idade:          req.Idade,
		Sexo:           req.Sexo,
		Peso:           req.Peso,
		Altura:         req.Altura,
		TipSanguineo:   req.TipSanguineo,
		GtsMusical:     req.GtsMusical,
		PraticaEsporte: req.PraticaQualEsporte,
		JogoPrefer:     req.JogoPrefer,
	}
	zombie := utils.ToZombie(hospedeiro)
	sang := strings.ToLower(req.TipSanguineo)
	if sang == "a+" || sang == "a-" || sang == "b+" || sang == "b-" || sang == "o+" || sang == "o-" || sang == "ab+" || sang == "ab-" {
		hospedeiro.TipSanguineo = req.TipSanguineo
	} else {
		ctx.JSON(400, gin.H{
			"erro": "tipo sanguíneo desconhecido",
		})
		return
	}

	if req.Nome != "" && req.Idade >= 0 && req.Sexo != "" && req.Peso > 0 && req.Altura > 0 && req.TipSanguineo != "" && req.GtsMusical != "" && req.PraticaQualEsporte != "" && req.JogoPrefer != "" {
		err := db.Create(&hospedeiro).Error
		if err != nil {
			ctx.JSON(400, gin.H{
				"erro": "erro ao tentar cadastrar hospedeiro",
			})
			return
		}
		err = db.Create(&zombie).Error
		if err != nil {
			ctx.JSON(400, gin.H{
				"erro": "erro db.Create(&zombie)",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"dados": hospedeiro,
			"analise": schemas.Analise{
				Forca:       zombie.Forca,
				Velocidade:  zombie.Velocidade,
				Resistencia: zombie.Resistencia,
				Classificao: zombie.Classificao,
			},
			"mensagem": "Hospedeiro cadastrado e analisado",
		})
	} else {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(400, gin.H{
			"erro": "Nenhuma informação deve ser nula",
		})
	}

}
