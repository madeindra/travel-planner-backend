package models

import "github.com/jinzhu/gorm"

type (
	UsersInterface interface {
		SelectAll() []Users
		SelectByID(id int) Users
	}

	UsersImplementation struct {
		db *gorm.DB
	}
)

func NewUsersImplementation(db *gorm.DB) *UsersImplementation {
	return &UsersImplementation{db}
}

func (impl *UsersImplementation) SelectAll() []Users {
	users := []Users{}
	impl.db.Find(&users)
	return users
}

func (impl *UsersImplementation) SelectByID(id int) Users {
	users := Users{}
	impl.db.Find(&users, id)
	return users
}
