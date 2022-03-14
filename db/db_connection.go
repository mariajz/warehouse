package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/tw-catalyst/dev-exercises/neev-13/pairs/pair-03/billfinder/config"
	"log"
)

func CreateDBConnection() {
	appConfig := config.New("./config.env")
	dbConfig := appConfig.DBConfig
	connectionStr := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DbPort, dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbName)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Printf("Error in creation of connection")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	fmt.Println("Connection successful")
	fmt.Println(db.Query("select * from bills"))
}
