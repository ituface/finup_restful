package main

import (
	fu "finup/controller"
	"fmt"
)

func main() {
	a:=fu.SqlSelectRequired("10003033")
	fmt.Println(a)

}
