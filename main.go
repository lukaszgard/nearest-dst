package main

import (
	"flag"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var logFilename, port string

func init() {
	flag.StringVar(&logFilename, "log-file", "", "name of logfile")
	flag.StringVar(&port, "port", "8080", "port application")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if len(logFilename) > 0 {
		f, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.SetOutput(os.Stdout)
			log.Panic("Unable to open/create file: ", logFilename)
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}
	log.SetLevel(log.TraceLevel)
}

func main() {
	log.Info("Starting Rest API")

	router := newRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
