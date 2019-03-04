package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func deleteAuth(writer http.ResponseWriter, request *http.Request) {

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

	rows, err := db.Query("select Id from Account where id =$1;", auth.ID)
	defer rows.Close()

	var result Resp
	if err != nil {
		//TODO sql delete
		//row, err := db.Exec("sql delete...")

		result.Code = "0"
		result.Msg = ""
		result.Result = "Isok:true"
	} else {
		result.Code = "402"
		result.Msg = "delete fail"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
	err = db.Close()
}
