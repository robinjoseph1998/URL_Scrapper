package routes

import (
	"URL_Scrapper/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {

	r.POST("/scrapper", controllers.UrlScrapper)

}
