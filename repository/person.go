package repository

import (
	"bogdanbarna/hello-rest-gin/config"
	"bogdanbarna/hello-rest-gin/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var persons []model.Person

type Repository struct {
	Database *gorm.DB
}

var db *gorm.DB
var err error

func FindPersons() ([]model.Person, error) {
	res := db.Find(&persons)
	return persons, res.Error
}

func FindPerson(username string) (model.Person, error) {
	var person model.Person
	res := db.Where("Username = ?", username).Find(&person)
	return person, res.Error
}

func CreatePerson(p model.Person) error {
	res := db.Select("Username", "Birthday").Create(&p)
	return res.Error
}

func UpdatePerson(p model.Person, birthday string) error {
	res := db.Model(&p).Update("Birthday", birthday)
	return res.Error
}

func SoftDeletePerson(p model.Person) error {
	res := db.Delete(&p)
	return res.Error
}

func init() {
	db, err = gorm.Open(postgres.Open(config.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&model.Person{})

	/*
		 	sqlDB, err := db.DB()
			if err != nil {
				log.Fatal(err)
			}

			// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
			sqlDB.SetMaxIdleConns(10)

			// SetMaxOpenConns sets the maximum number of open connections to the database.
			sqlDB.SetMaxOpenConns(100)

			// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
			sqlDB.SetConnMaxLifetime(time.Hour)
	*/
}
