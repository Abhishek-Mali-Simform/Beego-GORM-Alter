package services

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func checkColumn(models interface{}, name string) bool {
	db := NewDatabase()
	return db.Migrator().HasColumn(models, name)
}

func checkConstraint(models interface{}, constraint string) bool {
	db := NewDatabase()
	return db.Migrator().HasConstraint(models, constraint)
}

func AlterColumn(models interface{}, fields ...map[string]string) error {
	db := NewDatabase()
	for _, columns := range fields {
		for name, constraint := range columns {
			fmt.Println(name)
			fmt.Println(constraint)
			if checkColumn(models, name) {
				if constraint != "" {
					logs.Notice("Constraints To be Updated")
					if checkConstraint(models, constraint) {
						logs.Notice("Deleting Old Constraint")
						err := db.Migrator().DropConstraint(models, constraint)
						logs.Error(err)
					}
					logs.Notice("Updating New Constraint")
					return db.Migrator().CreateConstraint(models, constraint)
				}
			} else {
				logs.Notice("Adding New Column")
				return db.Migrator().AddColumn(models, name)
			}
		}
	}
	return errors.New("Something Went Wrong")
}
