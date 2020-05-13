package models

type (
	Login struct {
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password" validate:"required"`
	}
	Users struct {
		ID       int    `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
		Email    string `json:"email" gorm:"type:varchar(255);"`
		Username string `json:"username" gorm:"type:varchar(255);"`
		Name     string `json:"name" gorm:"type:varchar(255);"`
		Phone    string `json:"phone" gorm:"type:varchar(255);"`
		Password string `json:"password" gorm:"type:varchar(255);"`
		Image    string `json:"image" gorm:"type:varchar(255);"`
	}
)
