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
	id := c.Param("id")
	// id = id[1:]
	fmt.Printf("type of %T", id)
	fmt.Print("id is", id, "this is id")
	id_final, err := strconv.ParseInt(id, 10, 64)
	fmt.Printf("type of %T", id_final)

	if err != nil {
		fmt.Printf("cannot  converting id string to int")
	}
	medicine, err := daos.FetchMedicineById(int(id_final))

	if err != nil {
		fmt.Println("error getting value by ID")
		fmt.Println(err)
		c.IndentedJSON(http.StatusOK, gin.H{
			"Message": " cannot find medicibe by ID " + id,
		})
	} else {
		c.IndentedJSON(200, medicine)
	}
}

func AddMedicine(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var medicineData beans.Medicine
	if err:=c.ShouldBindJSON(&medicineData); err== nil {
		if er :=database.Db.Where("tablet_name = ?",medicineData.TabletName).Find(&medicineData); er!=nil{
			c.JSON(400,gin.H{
				"err":"Medicine name already exists",
			})
		}
			err:=database.Db.Create(&medicineData)
			c.JSON(200,gin.H{
				"data":medicineData,
				"Err":err,
			})
	}else{
		c.JSON(404,gin.H{"err":"error while marshalling data"})
	}
}

func DeleteMedicine(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id := c.Param("id")
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Printf("cannot  convert  string to int (id) ")
	}

	medicine, err := daos.DeleteMedicineFromDB(int(id_int))

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(400, "deleted failure")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": medicine,
		})
	}

}
