package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/madecanggih/travel-planner-backend/helpers"
	"github.com/madecanggih/travel-planner-backend/models"
	"github.com/madecanggih/travel-planner-backend/resources"
)

type UserHandler struct {
	do  models.UsersInterface
	vld helpers.ValidatorInterface
}

func NewUsersHandler(ui models.UsersInterface, vld helpers.ValidatorInterface) *UserHandler {
	return &UserHandler{ui, vld}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	jwtToken := c.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	users := h.do.SelectAll()
	data := []resources.UserData{}

	for i := range users {
		data = append(data, resources.UserData{ID: users[i].ID, Email: users[i].Email, Username: users[i].Username, Name: users[i].Name, Phone: users[i].Phone, Image: users[i].Image})
	}
	res := resources.UsersResponse{Status: true, Message: GeneralSuccessMessage, Data: data}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetOneUser(c echo.Context) error {
	jwtToken := c.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	id := c.Param("id")

	UserID, err := strconv.Atoi(id)
	if err != nil {
		res := setErrorResponse(UserNotFoundMessage)
		return c.JSON(http.StatusNotFound, res)
	}

	user := h.do.SelectByID(UserID)
	if user.ID == 0 {
		res := setErrorResponse(UserNotFoundMessage)
		return c.JSON(http.StatusNotFound, res)
	}

	data := resources.UserData{ID: user.ID, Email: user.Email, Username: user.Username, Name: user.Name, Phone: user.Phone, Image: user.Image}
	res := resources.UserResponse{Status: true, Message: GeneralSuccessMessage, Data: data}
	return c.JSON(http.StatusOK, res)
}
