package main

import (
	"json-app1/backend/controllers"
	"json-app1/backend/utils"

	"github.com/gin-gonic/gin"
)

func init(){
  utils.LoadEnvVariables() 
  utils.ConnectDB()
}

func handleHello(ctx *gin.Context){
  ctx.JSON(200, gin.H{
    "message":"hello go",
  })
}


func main(){
  r := gin.Default()

  r.GET("/", handleHello)
  r.GET("/posts", controllers.PostsGetAll)
  r.GET("/post/:id", controllers.PostGetOne)
  r.POST("/post", controllers.PostCreate)
  r.PUT("/post/:id", controllers.PostUpdate)
  r.DELETE("/post/:id", controllers.PostDelete)

  r.Run()
}
