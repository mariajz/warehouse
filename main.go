package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {

	http.HandleFunc("/rice", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			log.WithFields(log.Fields{"method": request.Method}).Debugf("Valid Method %s for post handler", request.Method)
			writer.WriteHeader(200)
		case "POST":
			log.WithFields(log.Fields{"method": request.Method}).Debugf("Valid Method %s for post handler", request.Method)
			writer.WriteHeader(200)
		default:
			log.WithFields(log.Fields{"method": request.Method}).Debugf("Method %s not supported", request.Method)

			writer.WriteHeader(404)
		}

		result, _ := httputil.DumpRequest(request, true)
		fmt.Printf("%s", result)
	})
}
