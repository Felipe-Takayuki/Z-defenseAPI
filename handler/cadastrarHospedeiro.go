package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/gin-gonic/gin"
)

func CatalogHospHandler(ctx *gin.Context) {
	req := HospedeiroReq{}
	ctx.BindJSON(&req)

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
	zombie := ToZombie(req)
	if req.Nome != "" && req.Idade >= 0 && req.Sexo != "" && req.Peso > 0 && req.Altura > 0 && req.TipSanguineo != "" && req.GtsMusical != "" && req.PraticaQualEsporte != "" && req.JogoPrefer != "" {
		db.Create(&hospedeiro)
		db.Create(&zombie)
        imc := zombie.Peso / (zombie.Altura * zombie.Altura)
		ctx.Header("Content-type", "application/json")
		ctx.JSON(200, gin.H{
			"dados":   hospedeiro,
			"analise" : schemas.Analise{
				Forca: zombie.Forca,
				Velocidade: zombie.Velocidade,
				Resistencia: zombie.Resistencia,
				Classificao: zombie.Classificao,
			},
			"imc" : imc,
			"mensagem": "Hospedeiro cadastrado e analisado",
		})
	} else {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(400, gin.H{
			"Error": "Nenhuma informação deve ser nula",
		})
	}

}
