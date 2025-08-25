package main

import (
	"github.com/gin-gonic/gin"
	"github.com/topengdev/svi_backend/controllers"
	"github.com/topengdev/svi_backend/initializers"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main(){

	r := gin.Default()

	// Create Article
	r.POST("/article", controllers.PostsCreate )

	// Update Article
	r.POST("/article/:id", controllers.PostUpdate)
	r.PATCH("/article/:id", controllers.PostUpdate)
	r.PUT("/article/:id", controllers.PostUpdate)

	// Delete Article
	r.DELETE("/article/:id", controllers.PostDelete)

	// Get Articles
	r.GET("/articles/:limit/:offset", controllers.PostsList)

	// Get Article
	r.GET("/article/:id", controllers.PostGetByID)


	r.Run()
}
