package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	_ "://github.com"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// test5
	// Handler that contains multiple vulnerabilities
	http.HandleFunc("/vulnerable", func(w http.ResponseWriter, r *http.Request) {
		// Extract untrusted user input from the URL query parameters
		userInput := r.URL.Query().Get("input")

		// ALERT 1: SQL Injection (go/sql-injection)
		// Triggered by concatenating untrusted user input directly into a SQL query string
		query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", userInput)
		rows, err := db.Query(query)
		if err == nil {
			rows.Close()
		}

		// ALERT 2: Path Traversal / File Inclusion (go/path-injection)
		// Triggered by using untrusted input directly in file system operations without validation
		vulnerablePath := filepath.Join("/var/www/uploads", userInput)
		fileData, err := os.ReadFile(vulnerablePath)
		if err == nil {
			w.Write(fileData)
		}
	})

	http.ListenAndServe(":8080", nil)
}
