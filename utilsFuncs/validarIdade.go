package utilsfuncs

import "github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"

func ValidarIdade( zombie schemas.Zombie) schemas.Zombie {
	Idade := zombie.Idade
	if Idade >= 0 && Idade <= 8 {
		zombie.Forca += 10
		zombie.Velocidade += 12
		zombie.Resistencia += 7
	}

	if Idade >= 9 && Idade <= 17 {
		zombie.Forca += 15
		zombie.Velocidade += 17
		zombie.Resistencia += 12
	}
	if Idade >= 18 && Idade <= 30 {
		zombie.Forca += 23
		zombie.Velocidade += 25
		zombie.Resistencia += 14
	}
	if Idade >= 30 && Idade <= 57 {
		zombie.Forca += 23
		zombie.Velocidade += 25
		zombie.Resistencia += 17
	}
	if Idade >= 58 && Idade <= 75 {
		zombie.Forca += 20
		zombie.Velocidade += 18
		zombie.Resistencia += 16
	}
	if Idade >= 76 && Idade < 87 {
		zombie.Forca += 16
		zombie.Velocidade += 14
		zombie.Resistencia += 14
	} else {
		zombie.Forca += 14
		zombie.Velocidade += 11
		zombie.Resistencia += 13
	}
	
	return zombie
}
