package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"sujith/tabRemBackend/beans"
	"sujith/tabRemBackend/daos"
	"sujith/tabRemBackend/resources/database"

	"github.com/gin-gonic/gin"
)

func GetMedicineDetails(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	medicine := daos.FetchMedicineDetails()

	c.IndentedJSON(http.StatusOK, medicine)
}

func GetCurrentMedicineDetail(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var id = c.Param("ID")
	var query []beans.Medicine
	database.Db.Where("id", id).Find(&query)
	fmt.Println(query)
	c.IndentedJSON(
		http.StatusOK, query)
}

func AddMedicine(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var medicineData beans.Medicine
	if err := c.ShouldBindJSON(&medicineData); err != nil {
		fmt.Print(medicineData)
	} else {
		fmt.Print("failed to add data to databse")
	}
	database.Db.Create(&medicineData)
}

func DeleteMedicine(c *gin.Context) {
	var id = c.Param("ID")
	id_int, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Printf("cannot  convert  string to int (id) ")
	}

	medicine, err := daos.DeleteMedicineFromDB(int(id_int))

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(400, "deleted failure")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": medicine.TabletName + " Deleted Succesfully",
		})
	}

}
