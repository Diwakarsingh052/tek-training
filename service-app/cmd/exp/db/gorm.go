package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type student struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique;not null"`
}

var db *gorm.DB

func main() {
	var err error
	l := logger.New(log.New(os.Stdout, "", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      logger.Info,
	})
	dsn := "host=localhost user=diwakar password=root dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: l})
	if err != nil {
		log.Fatalln(err)
	}
	//CreateTable()
	//InsertData()

	//SearchAll()
	SearchWhere()
}

func CreateTable() {
	// Drop the table student if it exists
	err := db.Migrator().DropTable(&student{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrator().AutoMigrate(student{})
	if err != nil {
		log.Fatalln(err)
	}

}

func InsertData() {
	s := []student{
		{
			Name:  "abc",
			Email: "diwakar@email.com",
		},
		{
			Name:  "abc",
			Email: "raj@email.com",
		},
		{
			Name:  "dev",
			Email: "dev@email.com",
		},
	}

	tx := db.Create(&s)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

}

func SearchAll() {
	var s []student
	err := db.Find(&s).Error
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
	// use db.find to fetch everything

}

// SearchWhere is a function that returns the record of a student where the name equals "Raj"
func SearchWhere() {
	var s []student
	name := "abc"

	tx := db.Where("name = ?", name)
	err := tx.Find(&s).Error
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}

// Delete is a function that deletes the first record in the student table
func Delete() {

	var s student
	// Fetch a first object found in DB to delete it later.
	db.First(&s)
	db.Delete(&s)
	// Unscoped method allows hard deleting an object
	// which means the record will be permanently deleted from the database
	// Unscoped is GORM's way of saying that it should ignore the soft delete
	err := db.Unscoped().Delete(&s).Error
	if err != nil {
		log.Println(err)
		return
	}
}
