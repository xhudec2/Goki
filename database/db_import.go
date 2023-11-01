package database

import (
	"log"

	"gorm.io/gorm"
)

func ImportDB[T Gettable](source *gorm.DB, target *gorm.DB, tableName string) error {
	var data []T
	if err := source.Table(tableName).Find(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	if err := target.Table(tableName).Create(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
