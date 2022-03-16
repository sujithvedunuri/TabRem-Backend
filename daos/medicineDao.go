package daos

import (
	"fmt"
	"log"
	"sujith/tabRemBackend/beans"
	"sujith/tabRemBackend/resources/database"

	"github.com/gin-gonic/gin"
)

//fetch all details from database
func FetchMedicineDetails() []beans.Medicine {
	var medicines []beans.Medicine
	database.Db.Find(&medicines)
	return medicines
}

//get details from database using id
func FetchMedicineById(id int) ([]beans.Medicine, error) {
	var medicine []beans.Medicine
	result := database.Db.Where("id = ?", id).Find(&medicine)
	return medicine, result.Error
}

//add details to database
func AddMedicineDetials(c *gin.Context) {
	var medicines beans.Medicine
	if err := c.ShouldBindJSON(&medicines); err == nil {
		database.Db.Create(&medicines)
	} else {
		log.Fatal(err)
	}
	meds := c.Params
	log.Fatal(meds)
	// database.Db.Create(&medicines)
}

//delete row from table
func DeleteMedicineFromDB(id int) (beans.Medicine, error) {
	var medicine beans.Medicine
	fmt.Println(id)
	result := database.Db.Where("id = ?", id).First(&medicine)
	if result.Error != nil {
		return medicine, result.Error
	} else {
		result := database.Db.Delete(&medicine, id)
		return medicine, result.Error

	}

}
