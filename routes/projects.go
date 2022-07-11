package routes

import (
	"log"
	"net/http"

	"github.com/Devansh3712/portfolio/api"
	"github.com/Devansh3712/portfolio/errors"
	"github.com/gin-gonic/gin"
)

// @Summary		Get projects of user
// @Produce 	json
// @Success 	200		{object} 	models.GitHub
// @Failure 	502		{object} 	errors.Error
// @Router 		/projects [get]
func Projects(ctx *gin.Context) {
	projects, err := api.GetRepos()
	if err != nil {
		log.Fatal(err)
		ctx.IndentedJSON(http.StatusBadGateway, errors.APIFetchingError)
		return
	}
	ctx.IndentedJSON(http.StatusOK, projects)
}
