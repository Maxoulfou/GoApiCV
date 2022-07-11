package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/Devansh3712/portfolio/config"
	"github.com/Devansh3712/portfolio/errors"
	"github.com/Devansh3712/portfolio/models"
	"github.com/gin-gonic/gin"
)

// @Summary		Get about section of user
// @Produce 	json
// @Success 	200		{object} 	models.About
// @Failure 	502		{object}	errors.Error
// @Router 		/about [get]
func About(ctx *gin.Context) {
	data := models.About{}
	if err := config.GetYaml("./resources/about.yaml", &data); err != nil {
		log.Fatal(err)
		ctx.IndentedJSON(http.StatusBadGateway, errors.DataParsingError)
		return
	}
	data.Timestamp = time.Now()
	ctx.IndentedJSON(http.StatusOK, data)
}
