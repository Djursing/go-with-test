package main

import (
	// "database/sql"
	// "fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	host   = "localhost"
	port   = 5432
	user   = "oliverdjursing"
	dbname = "test"
)

const dbFileName = "game.db.json"

func main() {
	// // connection string
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// // open DB
	// db, err := sql.Open("postgres", psqlconn)
	// CheckError(err)

	// // close DB
	// defer db.Close()

	// // check DB connection
	// err = db.Ping()
	// CheckError(err)

	// fmt.Println("Connected!")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	CheckError(err)

	store, err := poker.NewFileSystemPlayerStore(db)
	CheckError(err)

	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":9090", server))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
