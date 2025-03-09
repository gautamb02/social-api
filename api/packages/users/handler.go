package users

import (
	"log"
	"net/http"

	"github.com/gautamb02/social-api/api/rest"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) GetHTTPHandler() []*rest.HTTPHandler {
	return []*rest.HTTPHandler{
		{
			Version: 1,
			Method:  http.MethodGet,
			Path:    "users/{id}",
			Func:    h.GetUserByID,
		},
		{
			Version: 1,
			Method:  http.MethodPost,
			Path:    "users",
			Func:    h.CreateUser,
		},
	}
}

func (h *UserHandler) GetUserByID(c *rest.SessionContext) {
	user := h.userService.GetUserByID(1)
	c.Respond(http.StatusOK, true, "Success", user, "")
}

func (h *UserHandler) CreateUser(c *rest.SessionContext) {
	var req User
	err := c.BindBody(&req)

	if err != nil {
		log.Printf("Error: %s", err.Error())
		c.RespondWithError(http.StatusBadRequest, "Invalid Request Body")
		return
	}

	user, _ := h.userService.CreateUser(req)
	c.Respond(http.StatusOK, true, "Success", user, "")
}
