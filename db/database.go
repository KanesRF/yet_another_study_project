package db

import(
	"database/sql"
)

var connectString = "user=db_worker dbname=simple_db password=12345 host=localhost sslmode=disable"
var DbConn *sql.DB

func InitDB(){
	var err error
	DbConn, err = sql.Open("postgres", connectString)
	if err != nil {
  		panic(err)
	}
}

func CloseDB(){
	DbConn.Close()
}