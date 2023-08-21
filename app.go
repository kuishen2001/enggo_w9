package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/glebarez/go-sqlite"
)

// Define dbFile path
const dbDir = "db"
const dbFile = "QandA.db"

// App struct
type App struct {
	ctx       context.Context
	QueryInfo string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Search returns a query for the given name
func (a *App) Search(queryterm string) string {
	fmt.Println("Search query:", queryterm)
	return fmt.Sprintf("You are searching: %s", queryterm)
}

// Insert records into QandA.db before Query
func dbUpdate() error {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %s", err)
		return fmt.Errorf("error getting working directory: %s", err)
	}

	// Specify the full path to dbFile
	fn := filepath.Join(wd, dbDir, dbFile)
	//fn := "C:\\Users\\xxxxxxx\\Documents\\GitHub\\enggo\\enggo_w9\\db\\QandA.db"

	db, err := sql.Open("sqlite", fn)
	if err != nil {
		log.Fatalf("sql.Open failed: %s", err)
	}

	if _, err = db.Exec(`
INSERT OR IGNORE INTO qat (id, question, answer)
VALUES (24, "'default'", "'used within a switch statement to specify a block of code that should be executed when none of the case conditions match the value being switched'");
`); err != nil {
		return err
	}
	return nil
}

func (a *App) Query(queryterm string) string {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %s", err)
		return fmt.Sprintf("Error getting working directory: %s", err)
	}

	// Specify the full path to dbFile
	fn := filepath.Join(wd, dbDir, dbFile)

	db, err := sql.Open("sqlite", fn)
	if err != nil {
		log.Fatalf("sql.Open failed: %s", err)
	}

	// Call the dbUpdate function before executing the query
	if err := dbUpdate(); err != nil {
		log.Printf("Error updating database: %s", err)
		return "Error updating database: " + err.Error()
	}

	rows, err := db.Query("SELECT * FROM qat WHERE SUBSTR(question, 2, LENGTH(question) - 2) = ?", queryterm)
	if err != nil {
		log.Fatalf("db.Query failed: %s", err)
	}

	var queryResults string
	foundResults := false

	for rows.Next() {
		var id int
		var question string
		var answer string
		rows.Scan(&id, &question, &answer)
		queryResults += fmt.Sprintf("Your query for %s returned: %s\n", question, answer)
		foundResults = true
	}

	if !foundResults {
		queryResults = fmt.Sprintf("Your query for %s returned no result\n", queryterm)
	}

	a.QueryInfo = queryResults

	return a.QueryInfo // Return the query results and no error
}
