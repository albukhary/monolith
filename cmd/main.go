package main
import (
	"database/sql"
	"fmt"
	"github.com/albukhary/monolith/api/handlers"
	"github.com/albukhary/monolith/api"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbname = "crud"
	dbuser= "algorithmer"
	dbhost= "localhost"
	dbpass = "root"
	dbport = "5432"
	sslmode = "disable"
)

var connString = fmt.Sprintf(
	"host=%s, user=%s, password=%s, port =%s, dbname=%s, sslmode=%s",
	dbhost, dbuser, dbpass, dbport, dbname, sslmode)

func initDb(dbString string) *sql.DB {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatalln(err)
	}
 	return db
}


func main() {
	db := initDb(connString)
	defer db.Close()
	handler := handlers.NewHandler(db)
	app := api.New(handler)
	err := app.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}

