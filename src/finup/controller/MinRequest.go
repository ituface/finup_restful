package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// http.Get
func httpGet(url string)(s string,err error){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	body_s:=string(body)
	a:=strings.Split(body_s, "\"")
	return a[1],err
}

