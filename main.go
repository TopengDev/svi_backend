package main

import (
	"github.com/gin-gonic/gin"
	"github.com/topengdev/svi_backend/controllers"
	"github.com/topengdev/svi_backend/initializers"
	"github.com/gin-contrib/cors"
	"time"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main(){

	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())

  r.Use(cors.New(cors.Config{
    AllowOrigins: []string{
      "http://localhost:3000",
      "http://127.0.0.1:3000",
    },
    AllowMethods:     []string{"GET","POST","PUT","PATCH","DELETE","OPTIONS"},
    AllowHeaders:     []string{"Origin","Content-Type","Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: false, // set true only if you send cookies/auth headers
    MaxAge:           12 * time.Hour,
  }))

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

	// GET Deleted Articles
	r.GET("/articles/deleted/:limit/:offset", controllers.PostsListDeleted)


	// Get Article
	r.GET("/article/:id", controllers.PostGetByID)


	r.Run()
}
