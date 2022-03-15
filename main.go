package main

import (
	// "fmt"
	// "html/template"
	"fmt"
	"html/template"
	"log"
	"sujith/tabRemBackend/controllers"
	"sujith/tabRemBackend/resources/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {
	p := gin.Default()
	// googleProvider := google.New("128579797413-u853k52biqna8g7ht2jsiai4qbhv5ug0.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile")
	// goth.UseProviders(googleProvider)
	database.Init()

	p.GET("/getmedicines", controllers.GetMedicineDetails)
	p.GET("/getmedicines:id", controllers.GetCurrentMedicineDetail)
	p.POST("/addmedicine", controllers.AddMedicine)
	p.POST("/getUserDetails", controllers.GetUserDetails)
	// r.Run(":3000")

	key := securecookie.GenerateRandomKey(12) // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                      // 30 days
	isProd := false                           // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New("128579797413-u853k52biqna8g7ht2jsiai4qbhv5ug0.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	p.GET("/auth/{provider}/callback", func(c *gin.Context) {

		res := c.Writer
		req := c.Request

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(res, user)
	})

	p.GET("/auth/{provider}", func(c *gin.Context) {
		res := c.Writer
		req := c.Request
		gothic.BeginAuthHandler(res, req)
	})

	p.GET("/", func(c *gin.Context) {
		res := c.Writer
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(res, false)
	})
	log.Println("listening on localhost:8080")
	p.Run(":8080")

}
