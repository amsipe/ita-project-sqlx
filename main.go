package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var dbx *sqlx.DB

func init(){
	db, err := sql.Open("mysql", "itafinalproject:Kb55_K49-gMe@tcp(mysql6.gear.host)/itafinalproject")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	dbx = sqlx.NewDb(db,"mysql")
}
func main(){
	
	r := GetRoutes()
	r.Run(":5000")
	defer dbx.Close()

}