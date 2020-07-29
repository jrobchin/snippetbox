package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "user:password@tcp(:3308)/snippetbox?parseTime=true", "Data source name for the mysql DB.")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Close DB connection before the main() function exits
	defer db.Close()

	// Initialize application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Start new web server on port 4000
	infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

// Open a connection and ping DB, returns connection if there are no errors
func openDB(dsn string) (*sql.DB, error) {
	// THis only initializes the connection pool
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// This tests that communication to the server works
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
