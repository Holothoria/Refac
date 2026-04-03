package main

import (
        "log"
        "net/http"



)

func main() {
	db, err := initDB("chat.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	if err := createTables(db); err != nil {
		log.Fatal(err)
	}
	log.Println("Tables ready!")
	
	log.Println("Database connected!")
	log.Println("Server started on :5000")
	
	
	if err= http.ListenAndServe(":5000", nil); err!= nil {
		log.Fatal(err)
	}
}
  
			
	
		
