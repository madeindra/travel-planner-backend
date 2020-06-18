package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/madecanggih/travel-planner-backend/helpers"
	"github.com/madecanggih/travel-planner-backend/models"
	"github.com/madecanggih/travel-planner-backend/resources"
)

type LocationHandler struct {
	do  models.LocationsInterface
	vld helpers.ValidatorInterface
}

func NewLocationsHandler(ui models.LocationsInterface, vld helpers.ValidatorInterface) *LocationHandler {
	return &LocationHandler{ui, vld}
}

func (h *LocationHandler) GetAllLocations(c echo.Context) error {
	jwtToken := c.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	locations := h.do.SelectAll()
	data := []resources.LocationData{}

	for i := range locations {
		data = append(data, resources.LocationData{ID: locations[i].ID, Name: locations[i].Name, Description: locations[i].Description, Longitude: locations[i].Longitude, Latitude: locations[i].Latitude})
	}
	res := resources.LocationsResponse{Status: true, Message: GeneralSuccessMessage, Data: data}

	return c.JSON(http.StatusOK, res)
}

func (h *LocationHandler) GetOneLocation(c echo.Context) error {
	jwtToken := c.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	id := c.Param("id")

	LocationID, err := strconv.Atoi(id)
	if err != nil {
		res := setErrorResponse(NotFoundMessage)
		return c.JSON(http.StatusNotFound, res)
	}

	location := h.do.SelectByID(LocationID)
	if location.ID == 0 {
		res := setErrorResponse(NotFoundMessage)
		return c.JSON(http.StatusNotFound, res)
	}

	data := resources.LocationData{ID: location.ID, Name: location.Name, Description: location.Description, Longitude: location.Longitude, Latitude: location.Latitude}
	res := resources.LocationResponse{Status: true, Message: GeneralSuccessMessage, Data: data}
	return c.JSON(http.StatusOK, res)
}

func (h *LocationHandler) PostLocation(context echo.Context) error {
	jwtToken := context.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return context.JSON(http.StatusUnauthorized, res)
	}

	var location models.Locations
	if err := context.Bind(&location); err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return context.JSON(http.StatusInternalServerError, res)
	}

	data := h.do.Create(location)

	result := resources.LocationData{ID: data.ID, Name: data.Name, Description: data.Description, Longitude: data.Longitude, Latitude: data.Latitude}
	res := resources.LocationResponse{Status: true, Message: GeneralSuccessMessage, Data: result}
	return context.JSON(http.StatusCreated, res)
}

func (h *LocationHandler) PutLocation(context echo.Context) error {
	jwtToken := context.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return context.JSON(http.StatusUnauthorized, res)
	}

	id := context.Param("id")
	LocationID, err := strconv.Atoi(id)
	if err != nil {
		res := setErrorResponse(NotFoundMessage)
		return context.JSON(http.StatusNotFound, res)
	}

	var location models.Locations
	if err := context.Bind(&location); err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return context.JSON(http.StatusInternalServerError, res)
	}

	data := h.do.Update(LocationID, location)
	if location.ID == 0 {
		res := setErrorResponse(NotFoundMessage)
		return context.JSON(http.StatusNotFound, res)
	}

	result := resources.LocationData{ID: data.ID, Name: data.Name, Description: data.Description, Longitude: data.Longitude, Latitude: data.Latitude}
	res := resources.LocationResponse{Status: true, Message: GeneralSuccessMessage, Data: result}
	return context.JSON(http.StatusOK, res)
}

func (h *LocationHandler) DeleteLocation(context echo.Context) error {
	jwtToken := context.Request().Header.Get("Authorization")
	err := h.vld.ValidateToken(jwtToken)
	if err != nil {
		res := setErrorResponse(UnauthorizedMessage)
		return context.JSON(http.StatusUnauthorized, res)
	}

	id := context.Param("id")
	LocationId, err := strconv.Atoi(id)
	if err != nil {
		res := setErrorResponse(NotFoundMessage)
		return context.JSON(http.StatusNotFound, res)
	}

	check := h.do.SelectByID(LocationId)
	if check.ID == 0 {
		res := setErrorResponse(NotFoundMessage)
		return context.JSON(http.StatusNotFound, res)
	}

	data := h.do.Delete(check)
	result := resources.LocationData{ID: data.ID, Name: data.Name, Description: data.Description, Longitude: data.Longitude, Latitude: data.Latitude}
	res := resources.LocationResponse{Status: true, Message: GeneralSuccessMessage, Data: result}
	return context.JSON(http.StatusOK, res)
}
