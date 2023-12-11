package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	log.Print("Server started ... ")
	connectionString := os.Getenv("postgres_connection_string")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	defer db.Close()
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("k8s method recived = " + r.Method))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Recived method %s, from host %s", r.Method, r.Host)
		result, err := db.Prepare("insert into queries.queries(queries) values ($1)")
		if err != nil {
			log.Printf("Failed to insert query %d", err)
		}
		defer result.Close()

		if _, err := result.Exec(fmt.Sprintf("Recived method %s, from host %s", r.Method, r.Host)); err != nil {
			log.Fatalf("Error inserting query %s: %v", r.Method, err)
		}
		log.Print("insert done ")
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("ListenAndServe log error: %s", err)
	}
}
