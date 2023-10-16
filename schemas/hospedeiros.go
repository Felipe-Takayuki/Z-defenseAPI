package schemas

import (
	"time"

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

type HospedeiroResponse struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"createdAt"`
	Nome               string    `json:"nome"`
	Idade              int       `json:"idade"`
	Sexo               string    `json:"sexo"`
	Peso               float64   `json:"peso"`
	Altura             float64   `json:"altura"`
	TipSanguineo       string    `json:"tipoSanguineo"`
	GtsMusical         string    `json:"gostoMusical"`
	PraticaQualEsporte string    `json:"esporte"`
	JogoPrefer         string    `json:"jogoPreferido"`
}

type Zombie struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"createdAt"`
	Nome               string    `json:"nome"`
	Idade              int       `json:"idade"`
	Sexo               string    `json:"sexo"`
	Peso               float64   `json:"peso"`
	Altura             float64   `json:"altura"`
	TipSanguineo       string    `json:"tipoSanguineo"`
	GtsMusical         string    `json:"gostoMusical"`
	PraticaQualEsporte string    `json:"esporte"`
	JogoPrefer         string    `json:"jogoPreferido"`
	Forca              int64     `json:"forca"`
	Velocidade         int64     `json:"velocidade"`
	Inteligencia       int64     `json:"inteligencia"`
}
