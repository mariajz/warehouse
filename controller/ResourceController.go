package controller

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"warehouse/db"

	log "github.com/sirupsen/logrus"
)

// func PostData() {
// 	res := resource.Resource{
// 		Id:       4,
// 		ItemName: "rice",
// 		Quantity: 10,
// 		City:     "Chennai",
// 	}

// 	payloadBuf := new(bytes.Buffer)
// 	json.NewEncoder(payloadBuf).Encode(body)
// 	req, _ := http.NewRequest("POST", url, payloadBuf)

// 	client := &http.Client{}
// 	res, e := client.Do(req)
// 	if e != nil {
// 		return e
// 	}

// 	defer res.Body.Close()

// 	fmt.Println("response Status:", res.Status)
// 	// Print the body to the stdout
// 	io.Copy(os.Stdout, res.Body)

// }

func GetData() {

}
func Handle() {

	http.HandleFunc("/rice", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			db.SelectFromDb()
			log.WithFields(log.Fields{"method": request.Method}).Debugf("Valid Method %s for post handler", request.Method)
			writer.WriteHeader(200)
		case "POST":
			db.InsertIntoDb()

			log.WithFields(log.Fields{"method": request.Method}).Debugf("Valid Method %s for post handler", request.Method)
			writer.WriteHeader(201)
		default:
			log.WithFields(log.Fields{"method": request.Method}).Debugf("Method %s not supported", request.Method)

			writer.WriteHeader(404)
		}

		result, _ := httputil.DumpRequest(request, true)
		fmt.Printf("%s", result)

	})
}
