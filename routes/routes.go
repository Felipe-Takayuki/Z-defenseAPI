package routes

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/handler"
	"github.com/gin-gonic/gin"
)


func IniciarRotas(route *gin.Engine) {
	handler.IniciarHandler()
	app := route.Group("/api")
	{
		app.GET("/hospedeiros", handler.ListHospsHandler)
		app.POST("/catalogar", handler.CatalogHospHandler)
		app.GET("/analises", handler.ListToZombiesHandler) 
	}
}
