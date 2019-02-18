package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// Index handler present a welcome information
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Trace("------------------")
	log.Trace("Handle index routes")
	fmt.Fprintf(w, "This is index respose from nearest-dst/ Rest API.")
}

// NearestDest handler is used to obtain and return the nearest/fastest point from
// src to many dst's #TODO: fix this ugly info!
func nearestDest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Del("Date")
	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")

	log.Trace("---------------------")
	log.Trace("Handle routes handler")
	destination := make([]string, 0)
	r.ParseForm()
	src := r.FormValue("src")

	if !verifyCoordinates(src) {
		fmt.Fprintf(w, "Wrong 'src' coordinates (longitude,latitude): %s\n", src)
		return
	}

	for k, v := range r.Form {
		if k == "dst" {
			for _, i := range v {
				if !verifyCoordinates(i) {
					fmt.Fprintf(w, "Wrong 'dst' coordinates (longitude,latitude): %s\n", i)
					return
				}
			}
			destination = v
		}
	}

	if len(r.FormValue("src")) == 0 {
		log.Warn("Lack of source Form; 'src': ", r.FormValue("src"))
		fmt.Fprintf(w, "'src' not provided, please provide 'src' and atleast one 'dst'\n")
		return
	} else if len(r.FormValue("src")) > 0 {
		log.Trace("Provided Form 'src': ", r.FormValue("src"))
	}

	if len(destination) == 0 {
		log.Warn("Lack of destination/s Form; 'dst': ", destination)
		fmt.Fprintf(w, "'dst' not provided, atleast 1 should be provided\n")
		return
	} else if len(destination) > 0 {
		log.Trace("Provided Form/s 'dst': ", destination)
	}

	var Destinations []Destination

	//TODO: OSRM API mostly fails returning "Too Many Requests",
	// anyway other messages should be also handled.
	failOnToManyRequests := false
	outMsg := OSRMErr{}

	var wg sync.WaitGroup
	wg.Add(len(destination))

	go func() {
		for _, dst := range destination {
			rr, re := calculateDestination(src, dst)
			defer wg.Done()
			if re.Message == "Too Many Requests" {
				log.Warn("Response from OSRM API: 'Too Many Requests'")
				failOnToManyRequests = true
				outMsg = re
			}
			Destinations = append(Destinations, rr)
		}
	}()
	wg.Wait()

	if failOnToManyRequests {
		d, err := json.MarshalIndent(outMsg, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(d)
		return
	}
	if len(Destinations) > 1 {
		log.Trace("Unsorted routes: ", Destinations)
		sortDestinations(Destinations)
		log.Trace("Sorted routes over distance: ", Destinations)
	}

	destinationSortedRoutes := DestinationRoute{
		Source:             src,
		SortedDestinations: Destinations,
	}
	d, err := json.MarshalIndent(destinationSortedRoutes, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	log.Trace("Sucesfully finished request")
	w.Write(d)
	return
}
