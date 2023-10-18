package utils

import "github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"

func ValidarIdade(zombie schemas.Zombie) Pontos {
	Idade := zombie.Idade
	if Idade >= 0 && Idade <= 8 {
		zombie.Forca += 15
		zombie.Velocidade += 18
		zombie.Resistencia += 11
	}

	if Idade >= 9 && Idade <= 17 {
		zombie.Forca += 26
		zombie.Velocidade += 30
		zombie.Resistencia += 25
	}
	if Idade >= 18 && Idade <= 30 {
		zombie.Forca += 39
		zombie.Velocidade += 42
		zombie.Resistencia += 33
	}
	if Idade >= 30 && Idade <= 57 {
		zombie.Forca += 32
		zombie.Velocidade += 33
		zombie.Resistencia += 29
	}
	if Idade >= 58 && Idade <= 75 {
		zombie.Forca += 27
		zombie.Velocidade += 25
		zombie.Resistencia += 24
	}
	if Idade >= 76 && Idade < 87 {
		zombie.Forca += 22
		zombie.Velocidade += 18
		zombie.Resistencia += 19
	} 
	
	if Idade >  87{
		zombie.Forca += 14
		zombie.Velocidade += 5
		zombie.Resistencia += 12
	}

	return Pontos{
		Forca: zombie.Forca,
		Resistencia: zombie.Resistencia,
		Velocidade: zombie.Velocidade,
	}
}
