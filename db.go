package main

import (
        "database/sql"
        "fmt"
       _ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	  if err!= nil {
		  return nil, fmt.Errorf("failed to open database: %w", err)
	  }
	  
	  if err := db.Ping(); err != nil {
		  return nil, fmt.Errorf("failed to connect to db: %w", err)
	  }
	  
	  return db, nil
 }
 
 
 func createTables(db *sql.DB) error{
	 queries := []string{
		 
	 `CREATE TABLE IF NOT EXISTS users (
	    id               INTEGER       PRIMARY KEY     AUTOINCREMENT,
	    username         TEXT          NO NULL         UNIQUE,
	    created_at   DATETIME DEFAULT   CURRENT_TIMESTAMP
	    );`,
		  
		`CREATE TABLE IF NOT EXISTS rooms (
		  id             INTEGER        PRIMARY KEY     AUTOINCREMENT,
		  name           TEXT           NO NULL         UNIQUE,  
		  created_at   DATETIME DEFAULT   CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS messages (
		   id             INTEGER        PRIMARY KEY     AUTOINCREMENT,
		   room_id        INTEGER        NOT NULL,
		   user_id        INTEGER        NOT NULL,  
           content        TEXT           NOT NULL,
           created_at     DATETIME       DEFAULT         CURRENT_TIMESTAMP,
           FOREIGN KEY    (room_id)      REFERENCES      rooms(id)
           FOREIGN KEY    (user_id)      REFERENCES      users(id)
           );`,
	   }
           
         for _, query := range queries {
			 if _, err:= db.Exec(query); err != nil {
			   return fmt.Errorf("failed to create tables: %w", err)
		   }
	   }
		   return nil
	   }
   
   
