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
	Phnumber string

}
var db *gorm.DB

func main () {

	runServer()
	databaseConnection()
//add to database
//	currentUser := (User{Name:"Wane", Surname:"Wilson", Phnumber:"2234562"})
//	for i := 0; i < 10; i++{
//
//		db.NewRecord(currentUser)
//		db.Create(&currentUser)
//
//	}
//
//
//	changedUser := User{}
//
// update
//	db.First(&changedUser, 7)
//	changedUser.Name = "Volodya"
//	db.Save(&changedUser)
//
//
//
//delete
//
//	deleteUser  := User{}
//	db.First(&deleteUser, 2)
//	db.Delete(&deleteUser)

	//
	////output
	//users := []User{}
	//db.Debug().Find(&users)
	//fmt.Println(users)
}


func showAllUsers (c *gin.Context) {

	users := []User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)

}
func showUserWithID (c *gin.Context) {

	showUser := User{}
	db.First(&showUser, c.Param("id"))
	c.JSON(http.StatusOK, showUser)


}

func deleteUserWithID (c *gin.Context){

	deleteUser  := User{}
	db.First(&deleteUser, c.Param("id"))
	db.Delete(&deleteUser)
	c.JSON(http.StatusOK, deleteUser)


}

func createUser (c *gin.Context){
	newUser := (User{Name:c.Param("name"), Surname:c.Param("surname"), Phnumber:c.Param("phnumber")})
	db.NewRecord(newUser)
	db.Create(&newUser)

}

func runServer() {

	r := gin.Default()
	r.GET("/users", showAllUsers)
	r.GET("/users/:id", showUserWithID)
	r.DELETE("/users/:id", deleteUserWithID)
	r.POST("users/:name/:surname/:phnumber", createUser)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run(":3000")

}
func databaseConnection(){
	db, err := gorm.Open("postgres", "host=localhost dbname=my_first_databse sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
}