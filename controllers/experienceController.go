package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leilaApp/leila-api/services"
	"net/http"
)

func GetExperiences (context *gin.Context) {
	context.String(http.StatusOK,services.GetExperiences())
}
