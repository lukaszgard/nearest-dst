package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {
	log.Info("Starting Rest API")

	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
