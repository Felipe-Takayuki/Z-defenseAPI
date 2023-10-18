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
	gorm.Model
	Nome               string
	Idade              int
	Sexo               string
	Peso               float64
	Altura             float64
	TipSanguineo       string
	GtsMusical         string
	PraticaQualEsporte string
	JogoPrefer         string
	Forca              int64
	Velocidade         int64
	Resistencia        int64
	Classificao        string
}

type Analise struct {
	Forca              int64 
	Velocidade         int64
	Resistencia        int64
	Classificao        string
}