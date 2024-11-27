package controllers

import (
	"json-app1/backend/models"
	"json-app1/backend/utils"

	"github.com/gin-gonic/gin"
)

var body struct {
	Title string
	Body  string
}

func PostsGetAll(ctx *gin.Context) {
	var posts []models.Post
	utils.Db.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostGetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Post
	utils.Db.First(&post, id)

	ctx.JSON(200, gin.H{
		"posts": post,
	})
}

func PostCreate(ctx *gin.Context) {

	ctx.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := utils.Db.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "post is created",
		"post":    post,
	})
}

func PostUpdate(ctx *gin.Context) {
	id := ctx.Param("id")
  var post models.Post

	ctx.Bind(&body)

	utils.Db.First(&post, id)
	utils.Db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	ctx.JSON(200, gin.H{
		"message": "post is updated",
		"post":    body,
	})
}

func PostDelete(ctx *gin.Context){
  id := ctx.Param("id")
  utils.Db.Delete(&models.Post{}, id)

  ctx.Status(200)
}
