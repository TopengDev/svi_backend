package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/topengdev/svi_backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error

  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := os.Getenv("DB_URL")


  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

  log.Println("Connecting to database..... " + dsn)

  if err != nil{
	  log.Fatal("Failed to connect to Database: ", err)
  }

  fmt.Printf("DB.Logger: %v\n", DB.Logger)

  // auto-migrate
	if err := DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}


}
