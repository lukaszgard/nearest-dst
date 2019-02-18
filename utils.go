package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func calculateDestination(src, dst string) (Destination, OSRMErr) {
	osrm := OSRM{}
	osrmErr := OSRMErr{}
	route := Destination{}

	log.Trace("Obtaining route for dst: ", dst)
	response, err := http.Get("http://router.project-osrm.org/route/v1/driving/" + src + ";" + dst + "?overview=false")
	if err != nil {
		log.Fatal("The GET HTTP requests failed with error: ", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &osrm)
	if err != nil {
		if err != nil {
			log.Fatal("Failed with unmarshaling json: ", err)
		}
	}

	if osrm.Code == "Ok" {
		route = Destination{
			Destination: dst,
			Duration:    osrm.Routes[0].Duration,
			Distance:    osrm.Routes[0].Distance,
		}
	} else {
		err = json.Unmarshal(data, &osrmErr)
		if err != nil {
			log.Fatal("Failed with unmarshaling json: ", err)
		}
	}

	return route, osrmErr
}

func verifyCoordinates(coordinate string) bool {
	coordinates := strings.Split(coordinate, ",")
	if len(coordinates) == 2 {
		if checkCoordinateFormat(coordinates[0]) && checkCoordinateFormat(coordinates[1]) {
			return true
		}
	}
	return false
}

func checkCoordinateFormat(s string) bool {
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return true
	}
	return false
}
