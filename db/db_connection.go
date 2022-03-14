package db

import (
	"warehouse/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	fmt.Println(db.Query("select * from warehouses"))

	// sqlStatement := `
	// INSERT INTO warehouses (id, item_name, quantity, city)
	// VALUES ($1, $2, $3, $4)
	// RETURNING id`
	// id := 0
	// err = db.QueryRow(sqlStatement, 1, "rice", 10, "Mumbai").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)
}

