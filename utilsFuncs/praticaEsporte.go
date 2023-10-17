package utilsfuncs

import (
	"strings"

	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"
)

func PraticantedeEsporte(zombie schemas.Zombie) schemas.Zombie {
	esporte := strings.ToLower(zombie.PraticaQualEsporte)  
    if esporte == "esport" &&  esporte == "nenhum" {
		zombie.Forca -= 4 
		zombie.Resistencia -= 5
		zombie.Velocidade -= 4 
	}else {
		zombie.Forca +=10
		zombie.Velocidade += 8
		zombie.Resistencia += 10
	}
	return zombie
} 