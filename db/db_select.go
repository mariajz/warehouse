package db

import (
	"database/sql"
	"fmt"
	"log"
	"warehouse/config"
	resource "warehouse/model"

	_ "github.com/lib/pq"
)

func SelectFromDb() {
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

	rows, err := db.Query("SELECT * FROM warehouses WHERE  item_name= 'rice'")

	fmt.Println("Fetched records are:", rows)
	if err != nil {
		fmt.Errorf("error is %s", err)
		panic(err)
	}

	var result []resource.Resource

	for rows.Next() {
		var res resource.Resource
		if err := rows.Scan(&res.Id, &res.ItemName, &res.City, &res.Quantity); err != nil {
			fmt.Errorf("got error %v", err)
		}
		result = append(result, res)
	}
	fmt.Println(result)
	defer rows.Close()

}
