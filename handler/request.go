package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/utils"
)

type HospedeiroReq struct {
	Nome               string  `json:"nome"`
	Idade              int     `json:"idade"`
	Sexo               string  `json:"sexo"`
	Peso               float64 `json:"peso"`
	Altura             float64 `json:"altura"`
	TipSanguineo       string  `json:"tipo-sanguineo"`
	GtsMusical         string  `json:"gosto-musical"`
	PraticaQualEsporte string  `json:"pratica-esporte"`
	JogoPrefer         string  `json:"jogo-preferido"`
}


func ToZombie(hosp schemas.Hospedeiro) schemas.Zombie {
	zombie := schemas.Zombie{
		Nome:               hosp.Nome,
		Idade:              hosp.Idade,
		Sexo:               hosp.Sexo,
		Peso:               hosp.Peso,
		Altura:             hosp.Altura,
		TipSanguineo:       hosp.TipSanguineo,
		GtsMusical:         hosp.GtsMusical,
		PraticaEsporte: hosp.PraticaEsporte,
		JogoPrefer:         hosp.JogoPrefer,
	}

	zombie.Velocidade += utils.CalcImc(zombie, utils.Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Velocidade + utils.PraticantedeEsporte(zombie).Velocidade + utils.ValidarIdade(zombie).Velocidade
	zombie.Forca += utils.CalcImc(zombie, utils.Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Forca + utils.PraticantedeEsporte(zombie).Forca + utils.ValidarIdade(zombie).Forca
	zombie.Resistencia += utils.CalcImc(zombie, utils.Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Resistencia + utils.PraticantedeEsporte(zombie).Resistencia + utils.ValidarIdade(zombie).Resistencia
	zombie.Classificao = utils.ClassificarZombie(zombie)
    
	if (zombie.Velocidade > 100) {
		zombie.Velocidade = 100
	} 
	if (zombie.Forca > 100) {
		zombie.Forca = 100
	}
	if(zombie.Resistencia > 100) {
		zombie.Resistencia = 100 
	}
	if (zombie.Velocidade < 0) {
		zombie.Velocidade = 2
	}
	if (zombie.Forca < 0) {
		zombie.Forca = 2
	}
	if (zombie.Resistencia < 0){
		zombie.Resistencia = 2
	}
	return zombie

}
