package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/madecanggih/travel-planner-backend/helpers"
	"github.com/madecanggih/travel-planner-backend/models"
	"github.com/madecanggih/travel-planner-backend/resources"
)

type AuthHandler struct {
	do   models.AuthInterface
	help helpers.HelperInterface
}

func NewAuthHandler(ai models.AuthInterface, help helpers.HelperInterface) *AuthHandler {
	return &AuthHandler{ai, help}
}

func (h *AuthHandler) PostLogin(c echo.Context) error {
	var login models.Login

	if err := c.Bind(&login); err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}

	if login.Email == "" {
		res := setErrorResponse(EmailRequiredMessage)
		return c.JSON(http.StatusBadRequest, res)
	}

	if login.Password == "" {
		res := setErrorResponse(PasswordRequiredMessage)
		return c.JSON(http.StatusBadRequest, res)
	}

	userLogin := models.Users{Email: login.Email}
	users := h.do.SelectByEmail(userLogin)

	if users.ID == 0 {
		res := setErrorResponse(NotRegisteredMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	err := h.help.ComparePasswords(users.Password, login.Password)
	if err != nil {
		res := setErrorResponse(LoginErrorMessage)
		return c.JSON(http.StatusUnauthorized, res)
	}

	token, err := h.help.GenerateToken(users.Email, time.Now())
	if token == "" || err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}

	data := resources.LoginData{ID: users.ID, Email: users.Email, Username: users.Username, Name: users.Name, Phone: users.Phone, Image: users.Image, Token: token}
	res := resources.LoginResponse{Status: true, Message: LoginSuccessMessage, Data: data}
	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) PostRegister(c echo.Context) error {
	var register models.Users
	var hashedPassword string
	if err := c.Bind(&register); err != nil {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}

	hashedPassword = h.help.HashAndSalt(register.Password)
	if hashedPassword == "" {
		res := setErrorResponse(InternalServerErrorMessage)
		return c.JSON(http.StatusInternalServerError, res)
	}
	register.Password = hashedPassword

	check := h.do.NewUser(register)

	if check.ID != 0 {
		res := setErrorResponse(AlreadyRegisteredMessage)
		return c.JSON(http.StatusConflict, res)
	}

	users := h.do.Create(register)

	data := resources.UserData{ID: users.ID, Email: users.Email, Username: users.Username, Name: users.Name, Phone: users.Phone, Image: users.Image}
	res := resources.RegisterResponse{Status: true, Message: RegistrationSuccessMessage, Data: data}

	return c.JSON(http.StatusCreated, res)
}
