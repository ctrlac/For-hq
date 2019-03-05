package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createAuth(writer http.ResponseWriter, request *http.Request) {

	db, err := sql.Open("mysql", "root:a123@tcp(127.0.0.1:3306)/Test")

	// if err != nil {
	// 	fmt.Println("error")
	// } else {
	// 	fmt.Println("ok")
	// }

	var auth Auth
	if err := json.NewDecoder(request.Body).Decode(&auth); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}

	rows, err := db.Query("select Id from Account where id = $1;", auth.ID)
	defer rows.Close()

	var result Resp
	if err != nil {
		//rs, err := db.Exec("insert into Accoint value(Username=$1,Pwd=$2,id=$3;", auth.Username, auth.Pwd, auth.ID)
		result.Code = "0"
		result.Msg = ""
		result.Result = "Isok:true"
	} else {
		result.Code = "401"
		result.Msg = "create fail"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
	err = db.Close()
}
