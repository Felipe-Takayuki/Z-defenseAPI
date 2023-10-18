package utils

import "github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"

func ClassificarZombie(zombie schemas.Zombie) string {
	forca := zombie.Forca
	velocidade := zombie.Velocidade
	resistencia := zombie.Resistencia
	total :=( forca + velocidade + resistencia)/3 

	if total >= 74 && total <= 100 {
		zombie.Classificao = "Zumbi extremamente perigoso"
	}
	if total >= 62 && total <= 73{
		zombie.Classificao = "Zumbi perigoso"
	}
	if total >= 20 && total <= 62 {
		zombie.Classificao = "Zumbi de perigo normal"
	}
	if total < 20  {
		zombie.Classificao = "Zumbi fraco"
	}
	return zombie.Classificao
}
