package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-roundhouse"
	"github.com/whosonfirst/go-whosonfirst-roundhouse/http"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")
	var base = flag.String("base", "https://whosonfirst.mapzen.com/data/", "Where your Who's On First data lives")

	flag.Parse()

	rh := roundhouse.NewWOFRoundhouse()
	rh.Base = *base

	id_handler, err := http.IDHandler(rh)

	if err != nil {
		log.Fatal(err)
	}

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)

	mux := gohttp.NewServeMux()
	mux.Handle("/", id_handler)
	mux.Handle("/ping", ping_handler)	

	err = gohttp.ListenAndServe(address, mux)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
