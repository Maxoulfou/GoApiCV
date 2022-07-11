package errors

import "time"

type Error struct {
	Message   string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	MethodNotAllowed  = Error{Message: "Method not allowed.", Timestamp: time.Now()}
	EndpointNotFound  = Error{Message: "Endpoint not found.", Timestamp: time.Now()}
	DataParsingError  = Error{Message: "User data could not parsed.", Timestamp: time.Now()}
	APIFetchingError  = Error{Message: "Cannot fetch data from API.", Timestamp: time.Now()}
	MailParsingError  = Error{Message: "Invalid email.", Timestamp: time.Now()}
	SubjectEmptyError = Error{Message: "Email subject cannot be empty.", Timestamp: time.Now()}
	BodyEmptyError    = Error{Message: "Email body cannot be empty.", Timestamp: time.Now()}
	MailSendError     = Error{Message: "Could not send the email, try again later.", Timestamp: time.Now()}
	ServerError       = Error{Message: "An error occured while accessing the endpoint, try again later.", Timestamp: time.Now()}
)
