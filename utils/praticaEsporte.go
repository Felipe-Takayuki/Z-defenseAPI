package utils

import (
	"strings"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
)

func PraticantedeEsporte(zombie schemas.Zombie) Pontos {
	esporte := strings.ToLower(zombie.PraticaEsporte)  
    if esporte != "esport" &&  esporte != "nenhum" {
		zombie.Forca +=25
		zombie.Velocidade += 22
		zombie.Resistencia += 24
	}else {
		zombie.Forca +=4
		zombie.Velocidade += 5
		zombie.Resistencia += 2
	}
	return Pontos{
         Forca: zombie.Forca,
		 Resistencia: zombie.Resistencia,
		 Velocidade: zombie.Velocidade,
	}
} 