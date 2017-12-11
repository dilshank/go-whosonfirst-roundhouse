package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-roundhouse"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")
	var base = flag.String("base", "https://whosonfirst.mapzen.com/data/", "Where your Who's On First data lives")
	
	flag.Parse()

	rh := roundhouse.NewWOFRoundhouse()
	rh.Base = *base
	
	re := regexp.MustCompile("/([0-9]+)/?$")

	handler := func(rsp http.ResponseWriter, req *http.Request) {

		url := req.URL
		uri := url.RequestURI()

		m := re.FindStringSubmatch(uri)

		if len(m) == 0 {
			http.Error(rsp, "E_INSUFFICIENT_WOFID", http.StatusBadRequest)
			return
		}

		id := m[1]

		wofid, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			http.Error(rsp, "E_INVALID_WOFID", http.StatusBadRequest)
			return
		}

		u, err := rh.URL(wofid)

		if err != nil {
			http.Error(rsp, "E_IMPOSSIBLE_WOFID", http.StatusBadRequest)
			return
		}

		http.Redirect(rsp, req, u.String(), http.StatusSeeOther)
		return
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(endpoint, nil)

	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
