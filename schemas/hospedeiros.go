package schemas

import (
	"gorm.io/gorm"
)

type Hospedeiro struct {
	gorm.Model
	Nome           string
	Idade          int
	Sexo           string
	Peso           float64
	Altura         float64
	TipSanguineo   string
	GtsMusical     string
	PraticaEsporte string
	JogoPrefer     string
}

type Zombie struct {
	Nome               string  `json:"nome"`
	Idade              int     `json:"idade"`
	Sexo               string  `json:"sexo"`
	Peso               float64 `json:"peso"`
	Altura             float64 `json:"altura"`
	TipSanguineo       string  `json:"tipoSanguineo"`
	GtsMusical         string  `json:"gostoMusical"`
	PraticaQualEsporte string  `json:"esporte"`
	JogoPrefer         string  `json:"jogoPreferido"`
	Forca              int64   `json:"forca"`
	Velocidade         int64   `json:"velocidade"`
	Resistencia        int64   `json:"resistencia"`
	Classificao        string   `json:"classificação"`
}
