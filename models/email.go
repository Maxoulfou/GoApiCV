package models

type Email struct {
	Address string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}
