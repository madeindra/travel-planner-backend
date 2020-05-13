package helpers

import "time"

type (
	HelperInterface interface {
		ComparePasswords(hashedPwd, plainPwd string) error
		GenerateToken(email string, timenow time.Time) (string, error)
		HashAndSalt(pwd string) string
	}

	ValidatorInterface interface {
		ValidateToken(tokenString string) error
	}

	HelperImplementation struct{}

	ValidatorImplementation struct{}
)

func NewHelperImplementation() *HelperImplementation {
	return &HelperImplementation{}
}

func NewValidatorImplementation() *ValidatorImplementation {
	return &ValidatorImplementation{}
}

func (hi *HelperImplementation) ComparePasswords(hashedPwd, plainPwd string) error {
	return ComparePasswords(hashedPwd, plainPwd)
}

func (hi *HelperImplementation) GenerateToken(email string, timenow time.Time) (string, error) {
	return GenerateToken(email, timenow)
}

func (hi *HelperImplementation) HashAndSalt(pwd string) string {
	return HashAndSalt(pwd)
}

func (hi *ValidatorImplementation) ValidateToken(tokenString string) error {
	return ValidateToken(tokenString)
}
