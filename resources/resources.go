package resources

import "github.com/dgrijalva/jwt-go"

type (
	JwtClaims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}
	ErrorResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	RegisterResponse struct {
		Status  bool     `json:"status"`
		Message string   `json:"message"`
		Data    UserData `json:"data"`
	}

	LoginResponse struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    LoginData `json:"data"`
	}

	LoginData struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Image    string `json:"image"`
		Token    string `json:"token"`
	}

	UserData struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Image    string `json:"image"`
	}

	LocationData struct {
		ID             int                 `json:"id"`
		Name           string              `json:"name"`
		Description    string              `json:"description"`
		Longitude      string              `json:"longitude"`
		Latitude       string              `json:"latitude"`
		LocationImages []LocationImageData `json:"images"`
	}

	LocationImageData struct {
		ID     int    `json:"id"`
		Images string `json:"images"`
	}

	LocationResponse struct {
		Status  bool         `json:"status"`
		Message string       `json:"message"`
		Data    LocationData `json:"data"`
	}

	LocationsResponse struct {
		Status  bool           `json:"status"`
		Message string         `json:"message"`
		Data    []LocationData `json:"data"`
	}

	UserResponse struct {
		Status  bool     `json:"status"`
		Message string   `json:"message"`
		Data    UserData `json:"data"`
	}

	UsersResponse struct {
		Status  bool       `json:"status"`
		Message string     `json:"message"`
		Data    []UserData `json:"data"`
	}
)
