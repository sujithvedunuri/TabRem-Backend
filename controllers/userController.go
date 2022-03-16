package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sujith/tabRemBackend/beans"
	"sujith/tabRemBackend/resources/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

const Secretkey = "secret"

func LoginUser(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		log.Fatal("error")
	}

	var user beans.User

	database.Db.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(http.StatusNotFound)
		c.JSON(http.StatusNotFound, "user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.JSON(http.StatusUnauthorized, "password missmatched")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(Secretkey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// cookie := http.SetCookie(c.Writer, &http.Cookie{
	// 	Name:       "jwt",
	// 	Value:      token,
	// 	Path:       "",
	// 	Domain:     "",
	// 	Expires:    time.Now().Add(time.Hour * 24),
	// 	RawExpires: "",
	// 	MaxAge:     0,
	// 	Secure:     true,
	// 	HttpOnly:   true,
	// 	SameSite:   0,
	// 	Raw:        "",
	// 	Unparsed:   []string{},
	// })
	c.JSON(http.StatusOK, token)

}
func RegisterUser(c *gin.Context) {
	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(404, "cannot send data")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := beans.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}
	database.Db.Create(&user)
	c.JSON(http.StatusOK, user)

}
