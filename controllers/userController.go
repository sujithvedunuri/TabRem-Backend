package controllers

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println(string(jsonData))
	c.IndentedJSON(200, jsonData)

}
