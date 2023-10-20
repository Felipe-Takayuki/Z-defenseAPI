package utils

import "github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"

func ToZombie(hosp schemas.Hospedeiro) schemas.Zombie {
	zombie := schemas.Zombie{
		Nome:           hosp.Nome,
		Idade:          hosp.Idade,
		Sexo:           hosp.Sexo,
		Peso:           hosp.Peso,
		Altura:         hosp.Altura,
		TipSanguineo:   hosp.TipSanguineo,
		GtsMusical:     hosp.GtsMusical,
		PraticaEsporte: hosp.PraticaEsporte,
		JogoPrefer:     hosp.JogoPrefer,
	}

	zombie.Velocidade += CalcImc(zombie, Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Velocidade + PraticantedeEsporte(zombie).Velocidade + ValidarIdade(zombie).Velocidade
	zombie.Forca += CalcImc(zombie, Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Forca + PraticantedeEsporte(zombie).Forca + ValidarIdade(zombie).Forca
	zombie.Resistencia += CalcImc(zombie, Imc{Peso: zombie.Peso, Altura: zombie.Altura, Sexo: zombie.Sexo}).Resistencia + PraticantedeEsporte(zombie).Resistencia + ValidarIdade(zombie).Resistencia
	zombie.Classificao = ClassificarZombie(zombie)

	if zombie.Velocidade > 100 {
		zombie.Velocidade = 100
	}
	if zombie.Forca > 100 {
		zombie.Forca = 100
	}
	if zombie.Resistencia > 100 {
		zombie.Resistencia = 100
	}
	if zombie.Velocidade < 0 {
		zombie.Velocidade = 2
	}
	if zombie.Forca < 0 {
		zombie.Forca = 2
	}
	if zombie.Resistencia < 0 {
		zombie.Resistencia = 2
	}
	return zombie

}
