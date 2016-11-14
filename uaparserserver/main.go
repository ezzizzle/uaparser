package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ezzizzle/uaparser/useragent"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// UAResponses is a list of UAResponse structs
type UAResponses struct {
	Devices []useragent.Device
}

func parseUA(rw http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s IP: %s UA: %s", req.Method, req.RequestURI, req.RemoteAddr, req.UserAgent())
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	bodyString := string(bodyBytes)
	queryParams := req.URL.Query()

	uaResponses := UAResponses{}

	var userAgents []string

	// Get the user agents either from the post body or URL string
	if req.Method == "POST" {
		userAgents = strings.Split(bodyString, "\n")
	} else {
		// Look for the ua key in the query string
		if _, ok := queryParams["ua"]; ok {
			userAgents = queryParams["ua"]
		}
	}

	// TODO: Go routines or channels for this?
	for _, userAgent := range userAgents {
		if userAgent == "" {
			continue
		}
		// TODO: Implement caching of known UAs
		device := useragent.DeviceFromUA(userAgent)
		uaResponses.Devices = append(uaResponses.Devices, device)

	}

	// Pretty Print JSON?
	var jsonResponse []byte
	if _, ok := queryParams["pretty"]; ok {
		//do something here
		jsonResponse, _ = json.MarshalIndent(uaResponses, "", "    ")
	} else {
		jsonResponse, _ = json.Marshal(uaResponses)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonResponse)
}

func home(rw http.ResponseWriter, req *http.Request) {
	usageString := "POST /ua/parse - Post a list of user agents separated by newlines\nGET /ua/parse?ua=USERAGENT \nAdd ?pretty for pretty printed json"
	rw.Write([]byte(usageString))
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to use (Default 8080)")
	flag.Parse()

	portString := fmt.Sprintf(":%d", port)

	http.HandleFunc("/", home)
	http.HandleFunc("/ua/parse", parseUA)

	log.Printf("Listening on port %s", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}
