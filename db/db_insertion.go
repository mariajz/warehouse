package db

import (
	"database/sql"
	"fmt"
	"log"
	"warehouse/config"
	"warehouse/model"
)

func InsertIntoDb() {
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

	sqlStatement := `
	INSERT INTO warehouses (id, item_name, quantity, city)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	// id := 0

	res := resource.Resource{
		Id:       2,
		ItemName: "rice",
		Quantity: 10,
		City:     "Chennai",
	}

	err = db.QueryRow(sqlStatement, res.Id, res.ItemName,res.Quantity,res.City).Scan(&res.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", res.Id)
}
