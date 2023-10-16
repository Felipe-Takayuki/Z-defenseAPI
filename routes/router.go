package routes

import "github.com/gin-gonic/gin"

func Init() {
	//iniciando router
	router := gin.Default()

	//iniciar rotas
	IniciarRotas(router)
	//

	router.Run() // listen and serve on 0.0.0.0:8080
}
