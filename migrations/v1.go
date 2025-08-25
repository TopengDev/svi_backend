package main

import (
	"github.com/topengdev/svi_backend/initializers"
	"github.com/topengdev/svi_backend/models"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}
