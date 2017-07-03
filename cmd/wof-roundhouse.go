package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-roundhouse"
	"os"
	"strconv"
)

func main() {

	flag.Parse()

	r := roundhouse.NewWOFRoundhouse()

	for _, id := range flag.Args() {

		wofid, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			panic(err)
		}

		u, _ := r.URL(wofid)
		fmt.Println(u.String())
	}

	os.Exit(0)
}
