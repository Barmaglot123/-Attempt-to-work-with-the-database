package main

import (
	_ "github.com/lib/pq"
	_"github.com/jinzhu/gorm"

	//"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Car struct {
	Cid  int
	Manufacturer  string
	Model string

}


func main (){
	
	db, err := gorm.Open("postgres", "host=localhost dbname=myfisrtdb sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(db.Exec("SELECT * FROM \"cars\""))
	cars := make([]Car, 0)
	db.Find(&cars)
	fmt.Println(cars)



}
