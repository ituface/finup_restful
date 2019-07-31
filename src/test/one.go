package main

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
	a:=strings.Split(body_s, ",")
	fmt.Println(a)
	return a[2],err
}
func main() {

	body,_:=httpGet("http://10.10.180.206:8090/getEncrpyt?array=1821888,1212,1212")
	fmt.Println(body)
}