package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func login() {
	resp, _ := http.Get("/v1/user/login?username=admin&password=123456")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	fmt.Printf("Get request result: %s\n", string(body))
}
