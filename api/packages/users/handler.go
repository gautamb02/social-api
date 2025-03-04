package users

import (
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

func (h *UserHandler) GetUserByID(c *rest.SessionContext) {
	user := h.userService.GetUserByID(1)
	c.Respond(http.StatusOK, true, "Success", user, "")
}

func (h *UserHandler) GetHTTPHandler() []*rest.HTTPHandler {
	return []*rest.HTTPHandler{
		{
			Version: 1,
			Method:  http.MethodGet,
			Path:    "user/{id}",
			Func:    h.GetUserByID,
		},
	}
}
