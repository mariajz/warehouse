package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"warehouse/config"
	resource "warehouse/model"
	_ "github.com/lib/pq"
)

func InsertIntoDb(request *http.Request) {
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

	sqlStatement := `
	INSERT INTO warehouses (id, item_name, quantity, city)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	decoder := json.NewDecoder(request.Body)
	var res resource.Resource
	err = decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
	err = db.QueryRow(sqlStatement, res.Id, res.ItemName, res.Quantity, res.City).Scan(&res.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", res.Id)
}
