package models

import "github.com/jinzhu/gorm"

type (
	AuthInterface interface {
		SelectByEmail(data Users) Users
		Create(data Users) Users
		NewUser(data Users) Users
	}

	AuthImplementation struct {
		db *gorm.DB
	}
)

func NewAuthImplementation(db *gorm.DB) *AuthImplementation {
	return &AuthImplementation{db}
}

func (impl *AuthImplementation) SelectByEmail(data Users) Users {
	users := Users{}
	impl.db.Where(data).First(&users)
	return users
}

func (impl *AuthImplementation) Create(data Users) Users {
	impl.db.Create(&data)
	return data
}

func (impl *AuthImplementation) NewUser(data Users) Users {
	users := Users{Email: data.Email}
	impl.db.Where(users).First(&users)
	return users
}
