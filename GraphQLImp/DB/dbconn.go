package DB
//go env -w GO111MODULE=off
import (
	"log"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	//"os"
  )
  var db *gorm.DB

  func DBConn(){
	
	dsn := "postgres://zyjnpbwa:6kbWQ3f0S43kDBw2A6vkvmUAGlB-zhg5@kashin.db.elephantsql.com/zyjnpbwa"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	db.AutoMigrate(&Paciente{})

	if err != nil{ 
		log.Fatal("Failded to connect to database")
	}
	
  }

