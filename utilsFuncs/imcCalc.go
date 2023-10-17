package utilsfuncs

import (
	"strings"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
)

func CalcImc(zombie schemas.Zombie, imc Imc) schemas.Zombie {
	valueImc := imc.Peso / (imc.Altura * imc.Altura)

	if strings.ToLower(imc.Sexo) == "masculino" { //obesidade mórbida
		if valueImc > 43 {
			zombie.Forca -= 5
			zombie.Velocidade -= 25
			zombie.Resistencia += 5
		}
		if valueImc >= 30 && valueImc <= 39.9 { // obesidade moderada
			zombie.Forca -= 3
			zombie.Velocidade -= 10
			zombie.Resistencia += 6
		}
		if valueImc >= 25.9 && valueImc <= 29.9 { // sobrepeso
			zombie.Forca -= 2
			zombie.Velocidade -= 5
			zombie.Resistencia += 7
		}
		if valueImc >= 20 && valueImc <= 24.9 { // normal
			zombie.Forca += 5
			zombie.Velocidade += 10
			zombie.Resistencia += 14
		}
		if valueImc < 20 { //abaixo do peso
			zombie.Forca += 3
			zombie.Velocidade += 10
			zombie.Resistencia += 6
		}
	}
	
	if strings.ToLower(imc.Sexo) == "feminino" {
		if valueImc > 39 {
			zombie.Forca -= 5
			zombie.Velocidade -= 25
			zombie.Resistencia += 5
		}
		if valueImc >= 29 && valueImc <= 38.9 { // obesidade moderada
			zombie.Forca -= 3
			zombie.Velocidade -= 10
			zombie.Resistencia += 6
		}
		if valueImc >= 24.9 && valueImc <= 28.9 { // sobrepeso
			zombie.Forca -= 2
			zombie.Velocidade -= 5
			zombie.Resistencia += 7
		}
		if valueImc >= 19 && valueImc <= 23.9 { // normal
			zombie.Forca += 5
			zombie.Velocidade += 10
			zombie.Resistencia += 9
		}
		if valueImc < 19 { //abaixo do peso
			zombie.Forca += 3
			zombie.Velocidade += 10
			zombie.Resistencia += 6
		}
	}
	return zombie
	
}


type Imc struct {
	Peso   float64
	Altura float64
	Sexo   string
}
