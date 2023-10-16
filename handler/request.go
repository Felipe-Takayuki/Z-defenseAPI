package handler

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
