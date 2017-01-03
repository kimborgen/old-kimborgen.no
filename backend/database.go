package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db gorm.DB

//TODO delete this
var testing_password string

func dbStart() {
	dab, err := gorm.Open("postgres", "user=postgres password='postgres' dbname=kimborgen sslmode=disable")
	db = *dab
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}
	createTables()
}

func dbClose() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func createTables() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
	password, _ := bcrypt.GenerateFromPassword([]byte("kim"), bcrypt.DefaultCost)
	user := User{Name: "kim", Clearance: 9, Username: "kim", HashedPassword: string(password)}
	testing_password = user.HashedPassword
	db.Create(&user)
}

func destroyTables() {
	user := User{Name: "kim", Clearance: 9, Username: "kim"}
	db.Delete(&user)
}

/*

func migrate(kappa interface{}) {
	db.AutoMigrate(&kappa)
}

func create(kappa interface{}) {
	db.Create(&kappa)
}

func findAll(kappa interface{}) interface{} {
	db.Find(&kappa)
	return &kappa
}

func findSingle(kappa interface{}, id int) interface{} {
	db.First(&kappa, id)
	return &kappa
}

func update(kappa interface{}) {
	db.Save(&kappa)
}

func delete(kappa interface{}) {
	db.Delete(&kappa)
}

/*
func createTables() {
	if !production {
		files, _ := filepath.Glob("models/*")
		//var newFiles []string
		for _, file := range files {
			file = file[7 : len(file)-3]

			stmt, err := db.Prepare("SELECT to_regclass($1);")
			if err != nil {
				log.Fatal(err)
			}

			exist, err := stmt.Exec(file)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(exist)

		}
	}
}
*/
