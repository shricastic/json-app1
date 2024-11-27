package utils

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectDB() (){
  dsn := os.Getenv("DB_URL") 

  // dsn := "postgresql://postgre_owner:rGOhc9kuKb7D@ep-curly-dust-a16bkqok.ap-southeast-1.aws.neon.tech/postgre?sslmode=require"

  Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
  if err!=nil {
    log.Fatal("Error connecting DB:", err)
  } else{
    log.Println("Connected DB successfully!")
  } 
}
