package models

import (
	"GORM_BEEGO/services"
	"gorm.io/gorm"
)

type Student struct {
	Size           string `gorm:"check:size_checker,size <> '30'"`
	Gender         string
	MaratialStatus string
	Address        string
	Contact        string
	gorm.Model
	Id    int64 `gorm:"primaryKey"`
	Name  string
	Email string
}

func (student *Student) Insert() {
	db := services.NewDatabase()
	err := db.Create(&student).Error
	services.CheckErrorOrSuccess("Error Inserting Data", "Data Inserted Successfully", err)
}

func (student *Student) View() {
	db := services.NewDatabase()
	err := db.First(&student).Error
	services.CheckErrorOrSuccess("Data Not Exists", "Data Found Successfully", err)
}
