package handler

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
	utilsfuncs "github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/utilsFuncs"
)

type HospedeiroReq struct {
	Nome               string  `json:"nome"`
	Idade              int     `json:"idade"`
	Sexo               string  `json:"sexo"`
	Peso               float64 `json:"peso"`
	Altura             float64 `json:"altura"`
	TipSanguineo       string  `json:"tipoSanguineo"`
	GtsMusical         string  `json:"gostoMusical"`
	PraticaQualEsporte string  `json:"esporte"`
	JogoPrefer         string  `json:"jogoPreferido"`
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
		PraticaQualEsporte: hosp.PraticaEsporte,
		JogoPrefer:         hosp.JogoPrefer,
	}

	zombie.Velocidade += utilsfuncs.CalcImc(zombie, utilsfuncs.Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Velocidade + utilsfuncs.PraticantedeEsporte(zombie).Velocidade + utilsfuncs.ValidarIdade(zombie).Velocidade
	zombie.Forca += utilsfuncs.CalcImc(zombie, utilsfuncs.Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Forca + utilsfuncs.PraticantedeEsporte(zombie).Forca + utilsfuncs.ValidarIdade(zombie).Forca
	zombie.Resistencia += utilsfuncs.CalcImc(zombie, utilsfuncs.Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Resistencia + utilsfuncs.PraticantedeEsporte(zombie).Resistencia + utilsfuncs.ValidarIdade(zombie).Resistencia
	zombie.Classificao = utilsfuncs.ClassificarZombie(zombie).Classificao

	return zombie

}
