package handlers

import "github.com/madecanggih/travel-planner-backend/resources"

const GeneralSuccessMessage string = "Operation successful"
const BadRequestErrorMessage string = "Bad Request"
const NotFoundMessage string = "Not Found"
const InternalServerErrorMessage string = "Internal server error"
const RegistrationSuccessMessage string = "Registration successful"
const LoginSuccessMessage string = "Login successful"
const EmailRequiredMessage string = "Email is required"
const PasswordRequiredMessage string = "Password is required"
const NotRegisteredMessage string = "Email is not registered"
const LoginErrorMessage string = "Email or password is not correct"
const AlreadyRegisteredMessage string = "Email is already registered"
const UserNotFoundMessage string = "User not found"
const UnauthorizedMessage string = "Unauthorized"

func setErrorResponse(message string) resources.ErrorResponse {
	return resources.ErrorResponse{Status: false, Message: message}
}
