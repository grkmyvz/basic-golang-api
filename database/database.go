package database

import (
	"randgo/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Connect(config string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&utils.Login{},
		&utils.User{},
		&utils.Company{},
		&utils.CompanyService{},
		&utils.Appointment{},
		&utils.Comment{})
	if err != nil {
		panic(err)
	}

	return db
}
