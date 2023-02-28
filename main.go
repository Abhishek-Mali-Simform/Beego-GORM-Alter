package main

import (
	"GORM_BEEGO/models"
	_ "GORM_BEEGO/routers"
	"GORM_BEEGO/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//func init() {
//	db := services.NewDatabase()
//	err := db.AutoMigrate(&models.Student{})
//	services.CheckErrorOrSuccess("Error Migrating Model", "Model Migrated Successfully", err)
//}

func main() {
	//defer services.CloseDatabase()
	//student := models.Student{
	//	Name:           "Abhishek Mali",
	//	Email:          "abhishek.m@simformsolutions.com",
	//	Contact:        "9429865212",
	//	Address:        "Surat",
	//	MaratialStatus: "Single",
	//	Gender:         "Male",
	//	Size:           "28",
	//}
	err := services.AlterColumn(&models.Student{}, map[string]string{"size": "size_checker"})
	logs.Error(err)
	//student.Insert()
	//std := models.Student{Id: 1}
	//std.View()
	//logs.Info(std)
	beego.Run()
}
