package main

import (
	"fmt"
	"net/http"
	"os"
	"warehouse/config"
	"warehouse/db"
	controller "warehouse/controller"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	appConfig := config.New("./config.env")
	fmt.Println(appConfig)
	db.CreateDBConnection()

	controller.Handle()

	port := fmt.Sprintf(":%s", appConfig.ApplicationPort)
	_ = http.ListenAndServe(port, nil)
}

// body := &Student{
//     Name:    "abc",
//     Address: "xyz",
// }

// payloadBuf := new(bytes.Buffer)
// json.NewEncoder(payloadBuf).Encode(body)
// req, _ := http.NewRequest("POST", url, payloadBuf)

// client := &http.Client{}
// res, e := client.Do(req)
// if e != nil {
//     return e
// }

// defer res.Body.Close()

// fmt.Println("response Status:", res.Status)
// // Print the body to the stdout
// io.Copy(os.Stdout, res.Body)
