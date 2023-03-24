package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func insert(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "Method Now Allowed\n")
		return
	}

	param := InsertRequestParameter{}
	err := json.NewDecoder(req.Body).Decode(&param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, data := range param.Data {
		data.NewTimestamp, _ = time.Parse("2006-01-02 15:04:05", data.Timestamp)
		insertDatabase(data)
	}

}

func openDBConnection() {
	var err error
	connStr := "postgres://{USERNAME}:{PASSWORD}@localhost/{DATABASE_NAME}?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func insertDatabase(data InsertData) {
	// Insert DB
	_, err := db.Query("INSERT INTO transaction (id, customer, quantity, price, timestamp) VALUES ($1, $2, $3, $4, $5);", data.ID, data.Customer, data.Quantity, data.Price, data.NewTimestamp)
	if err != nil {
		log.Println("error insert data - " + err.Error())
	}
}

func main() {

	openDBConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/insert", insert)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
