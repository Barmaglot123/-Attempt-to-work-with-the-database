package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	"log"

)

type User struct {
	//Id int
	Name string
	Surname string
	Phone string
}
var db *gorm.DB

func main () {
	connectToDatabase()
	defer db.Close()
	runServer()
}


func usersList(c *gin.Context) {
	users := []User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
func userWithID (c *gin.Context) {
	user := User{}
	db.First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)

}

func deleteUserWithID (c *gin.Context){
	user  := User{}
	db.First(&user, c.Param("id"))
	db.Delete(&user)
	c.JSON(http.StatusOK, user)
}

func createUser (c *gin.Context){
	user := User{	Name:c.PostForm("name"),
			Surname:c.PostForm("surname"),
			Phone:c.PostForm("phnumber")}
	db.NewRecord(user)
	db.Create(&user)
	c.JSON(http.StatusOK, "Done!")


}

func runServer() {
	r := gin.Default()
	r.GET("/users", usersList)
	r.GET("/users/:id", userWithID)
	r.DELETE("/users/:id", deleteUserWithID)
	r.POST("users/:name/:surname/:phnumber", createUser)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run(":3000")
}
func connectToDatabase(){
	var err error
	db, err = gorm.Open("postgres", "host=localhost dbname=my_first_databse sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}