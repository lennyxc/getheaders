package main

import (
	"fmt"
    "net/http"
    "os"
    "log"
    "time"
)

func main() {
	//Set URL to test --- test will NOT follow redirects
	testurl := "https://beacon.walmart.com/vm/ttap.gif?id=10694084&160x600&audience=&creative=1027467&creativetype=rich_media&device=&initiative=209902&placement=SSLTEST&targeting=wmxaudience&vendor=AOD&version=2015802"

	// Set log filename
	logfile := "testlog.txt"

    var DefaultTransport http.RoundTripper = &http.Transport{}

    // Set number of requests
    for i := 0; i < 500000; i++ {

	    req, _ := http.NewRequest("GET", testurl, nil)

	    resp, _ := DefaultTransport.RoundTrip(req)

		logger, err := os.OpenFile(logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		defer logger.Close()

		// Grab specified headers from HTTP response
		location := resp.Header.Get("Location")
		cookie := resp.Header.Get("Set-Cookie")
		t := time.Now().UTC().String()
		result := fmt.Sprintf("%s\t%s\t%s\r\n", t, location, cookie)
		logger.WriteString(result)

		// Wait time before next request
		time.Sleep(time.Second)
	}
}