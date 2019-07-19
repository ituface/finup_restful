package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err:=sql.Open("mysql", "root:OYLDASuPfbpsEQB6@(10.10.231.135:3306)/lend_app")


	if err !=nil{
		fmt.Println("mysql connect is error")
	}
	defer db.Close()
	rows,_ :=db.Query(`select id,app_customer_id from app_lend_request where id=10002762`)
	fmt.Println(rows)
	var id,app_customer_id int
	for rows.Next(){
		rows.Scan(&id,&app_customer_id)
		fmt.Println(id,"---",app_customer_id)
	}
}
