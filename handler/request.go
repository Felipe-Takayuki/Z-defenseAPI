package handler


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


