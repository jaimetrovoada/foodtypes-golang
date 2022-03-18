package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func connectDb() *sql.DB {
	db, err := sql.Open("mysql", "jaime:jaimetdl@/foodtypes")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func queryData(db *sql.DB, foodName string) string {
	rows, err := db.Query("SELECT `SCIENTIFIC NAME` FROM `mockdata-food` WHERE `FOOD NAME` like ?", foodName)
	if err != nil {
		fmt.Println(err)
	}
	var name string
	for rows.Next() {
		var sciName string
		err = rows.Scan(&sciName)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println(sciName)
		name = sciName
	}

	return name
}

func startPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["name"])
	var foodName string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		foodName = v[0]
	}

	db := connectDb()
	scientificName := queryData(db, foodName)
	fmt.Fprintf(w, scientificName) // send data to client side
}

func startServer() {
	http.HandleFunc("/foodtypes", startPage) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
	// var foodName string
	// fmt.Print("Enter a food name: ")
	// fmt.Scanln(&foodName)
	// db := connectDb()
	// queryData(db, foodName)
}
