package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// Response struct to standardize API responses
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type SessionContext struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

// NewSessionContext initializes a new SessionContext
func NewSessionContext(r *http.Request, w http.ResponseWriter) *SessionContext {
	return &SessionContext{
		Request: r,
		Writer:  w,
	}
}

// Param extracts a URL parameter from the request
func (c *SessionContext) Param(key string) string {
	parts := strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/")
	for i, part := range parts {
		if part == key && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// BindBody binds the request body to the given object
func (c *SessionContext) BindBody(object interface{}) error {
	// Read the request body
	if err := render.DecodeJSON(c.Request.Body, object); err != nil {
		return err
	}

	// validationErrors, _ := validator.IsValid(obj)
	// if len(validationErrors) > 0 {
	// 	validationErrMsg := "falied to validate the request for fields"
	// 	for _, verr := range validationErrors {
	// 		validationErrMsg += fmt.Sprintf(" - %s", verr.Field)
	// 	}
	// 	return errors.New(validationErrMsg)
	// }
	return nil
}

// Respond sends a structured JSON response
func (c *SessionContext) Respond(statusCode int, success bool, message string, data interface{}, errMsg string) {
	response := Response{
		Success: success,
		Message: message,
		Data:    data,
		Error:   errMsg,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)

	if err := json.NewEncoder(c.Writer).Encode(response); err != nil {
		http.Error(c.Writer, `{"success": false, "error": "Failed to encode JSON"}`, http.StatusInternalServerError)
	}
}

// RespondWithError sends an error response with a custom message
func (c *SessionContext) RespondWithError(statusCode int, errMsg string) {
	c.Respond(statusCode, false, "", nil, errMsg)
}

// RespondWithData sends a success response with the given data
func (c *SessionContext) RespondWithData(statusCode int, message string, data interface{}) {
	c.Respond(statusCode, true, message, data, "")
}

// RespondWithMessage sends a success response with only a message
func (c *SessionContext) RespondWithMessage(statusCode int, message string) {
	c.Respond(statusCode, true, message, nil, "")
}
