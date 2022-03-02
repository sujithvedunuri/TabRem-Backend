package main

import (
	"sujith/tabRemBackend/controllers"
	"sujith/tabRemBackend/resources/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	database.Init()
	r.GET("/getMedicines", controllers.GetMedicineDetails)
	r.GET("/getMedicines:id", controllers.GetCurrentMedicineDetail)
	r.POST("/addMedicine", controllers.AddMedicine)
	r.Run()
}
