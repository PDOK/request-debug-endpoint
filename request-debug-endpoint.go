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
	p = flag.String("p", "80", "port number for de request-debug-endpoint")
	flag.Parse()
}

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprintf(w, `{"health": "OK"}`)
		return
	})

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {

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
			fmt.Fprintf(w, "\n =========== Request Body ===========\n %s \n =============== End ================", string(body))
			log.Printf("\n =========== Request Body ===========\n %s \n =============== End ================", string(body))
		}

		return
	})

	addr := ":" + *p
	log.Printf(`started on port: %s`, addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
