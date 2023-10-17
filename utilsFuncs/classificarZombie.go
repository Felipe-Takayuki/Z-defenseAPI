package utilsfuncs

import "github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/schemas"

func ClassificarZombie(zombie schemas.Zombie) schemas.Zombie {
   forca := zombie.Forca 
   velocidade := zombie.Velocidade 
   resistencia := zombie.Resistencia 
   total := forca + velocidade + resistencia   

	
   if total > 100 {
	zombie.Classificao = "Zumbi extremamente perigoso"
   }
   if total >= 75  &&  total <= 99 {
	zombie.Classificao = "Zumbi perigoso"
   }
   if total >= 45 && total <= 74 {
	zombie.Classificao = "Zumbi de perigo mediano"
   }
   if total >= 20 && total <= 44 {
	zombie.Classificao = "Zumbi de perigo normal"
   }
   if total >= 10 && total <= 19 {
	zombie.Classificao = "Zumbi de perigo abaixo do normal"
   } 
   if total <= 10{
	zombie.Classificao ="Zumbi fraco"
   }
   return zombie
}
