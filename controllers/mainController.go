package controllers

import (
	"fmt"
	"net/http"
	"sujith/tabRemBackend/daos"

	"github.com/gin-gonic/gin"
)

func GetMedicineDetails(c *gin.Context) {
	medicine := daos.FetchMedicineDetails()

	c.IndentedJSON(http.StatusOK, medicine)
}

func GetCurrentMedicineDetail(c *gin.Context) {
	fmt.Println("hello")
}

func AddMedicine(c *gin.Context) {
	fmt.Println("hello")
}
