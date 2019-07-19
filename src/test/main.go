package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"
)

var db *sql.DB

type Person struct {
	Id        int    `json:"id"`
	Product_name string `json:"first_name" form:"first_name"`
}
func (p Person) getAll() (persons []Person, err error) {
	rows, err := db.Query(`SELECT id, product_name FROM app_product`)
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.Product_name)
		persons = append(persons, person)
	}
	return

}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root123!@YE@(10.10.180.206:3306)/lend_app")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	router.GET("/persons", func(c *gin.Context) {
		var p Person
		persons, err := p.getAll()
		fmt.Println(persons)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})

	})

	router.Run(":8083") // listen and serve on 0.0.0.0:8083
}


