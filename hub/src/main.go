package main

import (
	"database/sql"
	"fmt"

	// "hub/src/database"
	"hub/src/database"
	"hub/src/handlers"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Main function that starts the server
func main() {
	// Get MySQL connection details from environment variables
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// Construct the MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	// Open the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open MySQL database:", err)
	}
	defer db.Close()

	// Check if the database connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to MySQL database:", err)
	}

	log.Println("Successfully connected to MySQL database")

	// A database handler to interact with the database. This is passed to the handlers
	dbHandler := database.NewDBHandler(db)

	// Create the subscribers table (if not already created)
	if err := database.CreateTable(db); err != nil {
		log.Fatal(err)
	}

	// Register the handlers for the root and generate paths
	http.HandleFunc("/", handlers.RootHandler(dbHandler))
	http.HandleFunc("/generate", handlers.GenerateHandler(dbHandler))

	http.ListenAndServe(":8080", nil)
}
