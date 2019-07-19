package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)


var My *sql.DB
func init() {
	var err error
	My,err=sql.Open("mysql", "root:OYLDASuPfbpsEQB6@(10.10.231.135:3306)/lend_app")


	if err !=nil{
		log.Fatalln("connect mysql errossr")
	}
	err=My.Ping()
	if err!=nil{
		log.Fatalln("ping is disconnect")
	}
//	defer My.Close()

}
