package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
)

type Car struct {
	Cid  int
	Manufacturer  string
	Model string

}


func main (){
	
	db, err := gorm.Open("postgres", "host=localhost dbname=myfisrtdb sslmode=disable")
	defer db.close()
	if err != nil {
		log.Fatal(err)
	}

	cars := new([]Car)
	db.Find(cars)
	fmt.Println(*cars)
}
