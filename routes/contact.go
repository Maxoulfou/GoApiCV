package routes

import (
	"net/http"
	"net/mail"
	"time"

	"github.com/Devansh3712/portfolio/api"
	"github.com/Devansh3712/portfolio/errors"
	"github.com/Devansh3712/portfolio/models"
	"github.com/gin-gonic/gin"
)

// @Summary		Send an email to the user
// @Accept		json
// @Produce 	json
// @Param		data 	body	models.Email	true	"Data for emailing"
// @Success 	200		{object} 	map[string]any
// @Failure		400		{object}	errors.Error
// @Failure 	502		{object} 	errors.Error
// @Router 		/contact [post]
func Contact(ctx *gin.Context) {
	var data models.Email
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.Error{
			Message: err.Error(), Timestamp: time.Now()},
		)
		return
	}
	if _, err := mail.ParseAddress(data.Address); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.MailParsingError)
		return
	}
	if data.Subject == "" {
		ctx.IndentedJSON(http.StatusBadRequest, errors.SubjectEmptyError)
		return
	}
	if data.Body == "" {
		ctx.IndentedJSON(http.StatusBadRequest, errors.BodyEmptyError)
		return
	}
	if err := api.SendEmail(data); err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, errors.MailSendError)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message":   "Email sent successfully! I will reply to you as soon as possible :)",
		"timestamp": time.Now(),
	})
}
