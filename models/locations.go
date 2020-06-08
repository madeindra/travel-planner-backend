package models

import "github.com/jinzhu/gorm"

type (
	LocationsInterface interface {
		SelectAll() []Locations
		SelectByID(id int) Locations
		Create(data Locations) Locations
		Update(id int, data Locations) Locations
		Delete(data Locations) Locations
	}

	LocationsImplementation struct {
		db *gorm.DB
	}
)

func NewLocationsImplementation(db *gorm.DB) *LocationsImplementation {
	return &LocationsImplementation{db}
}

func (impl *LocationsImplementation) SelectAll() []Locations {
	locations := []Locations{}
	impl.db.Find(&locations)
	return locations
}

func (impl *LocationsImplementation) SelectByID(id int) Locations {
	locations := Locations{}
	impl.db.Find(&locations, id)
	return locations
}

func (impl *LocationsImplementation) Create(data Locations) Locations {
	impl.db.Create(&data)
	return data
}

func (impl *LocationsImplementation) Update(id int, data Locations) Locations {
	locations := Locations{}
	impl.db.Model(&data).UpdateColumns(data).Find(&locations)
	return locations
}

func (impl *LocationsImplementation) Delete(data Locations) Locations {
	impl.db.Delete(&data)
	return data
}
