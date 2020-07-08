package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var p *string

func init() {
	// Get port if none provided use default :80
	p = flag.String("p", "80", "port number for de request-debug-endpoint")
	flag.Parse()
}

func main() {

	// Health (and sanity check endpoint)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprintf(w, `{"health": "OK"}`)
		return
	})

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {

		// Log
		log.Printf("Method: %v\n", r.Method)
		log.Printf("Host: %s\n", r.Host)
		log.Printf("Referer: %s\n", r.Referer())

		// Reflect
		fmt.Fprintf(w, "Method: %v\n", r.Method)
		fmt.Fprintf(w, "Host: %s\n", r.Host)
		fmt.Fprintf(w, "Referer: %s\n", r.Referer())

		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
				log.Printf("%v: %v\n", name, h)
			}
		}

		body, _ := ioutil.ReadAll(r.Body)
		if len(body) != 0 {
			str := string(body)
			fmt.Fprintf(w, "\n =========== Request Body ===========\n %s \n =============== End ================", str)
			log.Printf("\n =========== Request Body ===========\n %s \n =============== End ================", str)
		}

		return
	})

	addr := ":" + *p
	log.Printf(`started on port: %s`, addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
