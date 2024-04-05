package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "postgresql://store_owner:Ne4vQbT3aRjC@ep-twilight-sound-a1kbw9sd.ap-southeast-1.aws.neon.tech/gin?sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database")
	} else {
		log.Println("Database connection was made successfully")
	}
}
