package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
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
	zombie := ToZombie(hospedeiro)
	if req.Nome != "" && req.Idade >= 0 && req.Sexo != "" && req.Peso > 0 && req.Altura > 0 && req.TipSanguineo != "" && req.GtsMusical != "" && req.PraticaQualEsporte != "" && req.JogoPrefer != "" {
		err := db.Create(&hospedeiro).Error; if err != nil {
            ctx.JSON(400, gin.H{
				"erro" : "erro db.Create(&hospedeiro)",
			})
			return 
		}
		err = db.Create(&zombie).Error; if err != nil {
			ctx.JSON(400, gin.H{
				"erro" : "erro db.Create(&zombie)",
			})
			return 
		}

		
		ctx.JSON(200, gin.H{
			"dados":   hospedeiro,
			"analise" : schemas.Analise{
				Forca: zombie.Forca,
				Velocidade: zombie.Velocidade,
				Resistencia: zombie.Resistencia,
				Classificao: zombie.Classificao,
			},
			"mensagem": "Hospedeiro cadastrado e analisado",
		})
	} else {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(400, gin.H{
			"Error": "Nenhuma informação deve ser nula",
		})
	}

}
