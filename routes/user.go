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

// Get current age of user.
func getAge(year int, month int, date int) int {
	dob := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
	curr := time.Now()
	days := curr.Sub(dob).Hours() / 24
	return int(days / 365)
}

// Format user data.
func getUserData() (*models.User, error) {
	data := models.User{}
	if err := config.GetYaml("./resources/root.yaml", &data); err != nil {
		return nil, err
	}
	data.Age = getAge(data.Year, data.Month, data.Date)
	data.Timestamp = time.Now()
	return &data, nil
}

// @Summary		Get user information
// @Produce 	json
// @Success 	200		{object} 	models.User
// @Failure 	502		{object}	errors.Error
// @Router 		/ [get]
func Root(ctx *gin.Context) {
	userData, err := getUserData()
	if err != nil {
		log.Fatal(err)
		ctx.IndentedJSON(http.StatusBadGateway, errors.DataParsingError)
		return
	}
	ctx.IndentedJSON(http.StatusOK, userData)
}
