package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Resp struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Result string `json:"result"`
}

type Auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
	ID       string `json:"id"`
}

type apitype struct {
	create string
	delete string
	change string
	login  string
}

func main() {
	var Apitype apitype
	Apitype.create = "/v1/user/create"
	Apitype.delete = "/v1/user/delete"
	Apitype.change = "/v1/user/change"
	Apitype.login = "/v1/user/login"

	postWithJson(Apitype.create)
	postWithJson(Apitype.delete)
	postWithJson(Apitype.change)
	login()

}

func postWithJson(ApiType string) {
	
	auths := Auth{"admin", "123456", "A123456789"}
	ba, _ := json.Marshal(auths)
	resp, _ := http.Post(ApiType, "application/json", bytes.NewBuffer([]byte(ba)))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Post request with json result: %s\n", string(body))
}

func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}
