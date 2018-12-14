package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leilaApp/leila-api/controllers"
)

func AddRoutesExperiences(r *gin.RouterGroup) {
	r.GET("/experiences", controllers.GetExperiences)

}
