package controllers

import (
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

const Secretkey = "secret"

func LoginUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

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

	cookie1 := &http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Expires: time.Now().Add(time.Hour * 24)}
	http.SetCookie(c.Writer, cookie1)

	c.JSON(http.StatusOK, token)

}
func RegisterUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(404, "cannot send data")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := beans.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.Db.Create(&user)
	c.JSON(http.StatusOK, user)

}

func GetUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, er := c.Cookie("jwt")
	if er != nil {
		c.JSON(404, "error sujith")
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})
	if err != nil {
		c.JSON(401, err)
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user beans.User
	database.Db.Where("id =?", claims.Issuer).First(&user)


	c.IndentedJSON(http.StatusOK, user)
}
func Logout(c *gin.Context) {
	cookie1 := &http.Cookie{Name: "jwt", Value: "", HttpOnly: true, Expires: time.Now().Add(-time.Hour)}
	http.SetCookie(c.Writer, cookie1)
	c.JSON(200, "success")
}
