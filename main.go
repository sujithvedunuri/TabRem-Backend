package main

import (
	"log"
	"sujith/tabRemBackend/controllers"
	"sujith/tabRemBackend/resources/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// googleProvider := google.New("128579797413-u853k52biqna8g7ht2jsiai4qbhv5ug0.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile")
	// goth.UseProviders(googleProvider)
	database.Init()

	r.GET("/getmedicines", controllers.GetMedicineDetails)
	r.GET("/getmedicines/:id", controllers.GetCurrentMedicineDetail)
	r.POST("/addmedicine", controllers.AddMedicine)
	r.POST("/deletemedicine/:id", controllers.DeleteMedicine)
	r.POST("/api/login", controllers.LoginUser)
	r.POST("/api/register", controllers.RegisterUser)
	r.GET("/api/user", controllers.GetUser)
	r.POST("api/logout", controllers.Logout)

	// r.Run(":3000")
	log.Println("listening on localhost:8080")
	r.Run(":8080")

}
