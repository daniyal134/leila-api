package routers

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {


	
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.CORS())
	v1 := r.Group("api/v1")

	AddRoutesExperiences(v1)


	//// Rotas autorizadas: Adicionar aqui as rotas autorizadas
	//AddRoutesAuthentication(v1)
	//
	//v1.Use(middleware.AuthRequired())
	//{
	//	// Rotas não autorizadas: Adicionar aqui as rotas não autorizadas
	//	AddRoutesPeople(v1)
	//	AddRoutesTaskCategories(v1)
	//	AddRoutesTasks(v1)
	//}

	return r
}
