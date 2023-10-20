package schemas

import (
	"gorm.io/gorm"
)

type Hospedeiro struct {
	gorm.Model
	Nome           string  `json:"nome"`
	Idade          int     `json:"idade"`
	Sexo           string  `json:"sexo"`
	Peso           float64 `json:"peso"`
	Altura         float64 `json:"altura"`
	TipSanguineo   string  `json:"tipo-sanguineo"`
	GtsMusical     string  `json:"gosto-musical"`
	PraticaEsporte string  `json:"pratica-esporte"`
	JogoPrefer     string  `json:"jogo-preferido"`
}

type Zombie struct {
	gorm.Model
	Nome           string  `json:"nome"`
	Idade          int     `json:"idade"`
	Sexo           string  `json:"sexo"`
	Peso           float64 `json:"peso"`
	Altura         float64 `json:"altura"`
	TipSanguineo   string  `json:"tipo-sanguineo"`
	GtsMusical     string  `json:"gosto-musical"`
	PraticaEsporte string  `json:"pratica-esporte"`
	JogoPrefer     string  `json:"jogo-preferido"`
	Forca          int64   `json:"forca"`
	Velocidade     int64   `json:"velocidade"`
	Resistencia    int64   `json:"resistência"`
	Classificao    string  `json:"classificacao"`
}

type Analise struct {
	Forca       int64
	Velocidade  int64
	Resistencia int64
	Classificao string
}
