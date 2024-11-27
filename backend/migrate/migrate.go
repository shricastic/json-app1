package main

import (
	"json-app1/backend/models"
	"json-app1/backend/utils"
)

func init(){
  utils.LoadEnvVariables()
  utils.ConnectDB()
}

func main(){
  utils.Db.AutoMigrate(&models.Post{})
}
