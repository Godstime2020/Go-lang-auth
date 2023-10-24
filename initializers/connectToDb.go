package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  var DB *gorm.DB

  func ConnectToDb (){
	// postgres://ocftyaty:Sr1fXA77VjR5Y-4kHL23MqEfteSsmk42@mahmud.db.elephantsql.com/ocftyaty
	var err error

	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}
  }
  
