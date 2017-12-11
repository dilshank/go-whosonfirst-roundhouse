package main

// everything in here needs to be updated from moving the HTTP handler code
// in to the http package to https://github.com/whosonfirst/go-whosonfirst-roundhouse/issues/1
// (20171211/thisisaaronland)

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-crawl"
	"github.com/whosonfirst/go-whosonfirst-csv"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Repo struct {
	Index int
	Path  string
}

type Path struct {
	Id   int64
	Repo int
}

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")

	var prefix = flag.String("prefix", "", "...")
	var strict = flag.Bool("strict", false, "...")

	var max_fh = flag.Int("max-filehandles", 100, "The maximum number of filehandles to keep open while indexings repos")

	flag.Parse()

	paths := flag.Args()
	count_paths := len(paths)

	repo_map := make(map[int]string)
	lookup_map := make(map[int64]int)

	indexing := true

	indexing_ch := make(chan bool)
	repo_ch := make(chan Repo)
	path_ch := make(chan Path)

	fh_throttle := make(chan bool, *max_fh)

	for i := 0; i < *max_fh; i++ {
		fh_throttle <- true
	}

	pairs := 0

	go func() {

		for {

			select {
			case <-indexing_ch:
				indexing = false
				log.Printf("indexed %d pairs\n", pairs)
				break
			case r := <-repo_ch:
				repo_map[r.Index] = r.Path
			case p := <-path_ch:
				lookup_map[p.Id] = p.Repo
				pairs += 1
			}
		}

	}()

	go func(paths []string) {

		done_ch := make(chan bool)

		t1 := time.Now()

		for r, root := range paths {

			go func(root string, r int, repo_ch chan Repo, done_ch chan bool) {

				defer func() {
					done_ch <- true
				}()

				repo := filepath.Base(root)
				meta := filepath.Join(root, "meta")

				_, err := os.Stat(meta)

				if os.IsNotExist(err) {
					log.Printf("%s is missing meta directory\n", repo)

					if *strict {
						log.Fatal(err)
					}
				}

				repo_ch <- Repo{r, root}

				ta := time.Now()
				log.Printf("start indexing %s at %v", repo, ta)

				callback := func(path string, info os.FileInfo) error {

					if info.IsDir() {
						return nil
					}

					if !strings.HasSuffix(path, "-latest.csv") {
						return nil
					}

					<-fh_throttle

					defer func() {
						fh_throttle <- true
					}()

					reader, err := csv.NewDictReaderFromPath(path)

					if err != nil {
						return err
					}

					for {
						row, err := reader.Read()

						if err == io.EOF {
							break
						}

						if err != nil {
							return err
						}

						str_id, ok := row["id"]

						if !ok {
							continue
						}

						wofid, err := strconv.ParseInt(str_id, 10, 64)

						if err != nil {
							return err
						}

						// test that file exists here?

						path_ch <- Path{wofid, r}
					}

					return nil
				}

				c := crawl.NewCrawler(meta)
				err = c.Crawl(callback)

				if err != nil {
					log.Fatal()
				}

				tb := time.Since(ta)

				log.Printf("time to index %s: %v\n", repo, tb)

			}(root, r, repo_ch, done_ch)
		}

		for i := count_paths; i > 0; {
			select {
			case <-done_ch:
				i--
			}
		}

		t2 := time.Since(t1)
		log.Printf("time to index all: %v\n", t2)

		indexing_ch <- true

	}(paths)

	handler := func(rsp http.ResponseWriter, req *http.Request) {

		if indexing == true {
			http.Error(rsp, "indexing records", http.StatusServiceUnavailable)
			return
		}

		url := req.URL
		path := url.Path

		if *prefix != "" {
			path = strings.Replace(path, *prefix, "", 1)
		}

		is_wof, err := uri.IsWOFFile(path)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if !is_wof {
			http.Error(rsp, "Insufficient Who's On First", http.StatusNotFound)
			return
		}

		id, err := uri.IdFromPath(path)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		repo_idx, ok := lookup_map[id]

		if !ok {
			http.Error(rsp, "What is the what (lookup_map)", http.StatusNotFound)
			return
		}

		repo, ok := repo_map[repo_idx]

		if !ok {
			http.Error(rsp, "What is the what (repo_map)", http.StatusNotFound)
			return
		}

		fname := filepath.Base(path)
		parent, err := uri.Id2Path(id)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		root := filepath.Join(repo, "data")

		rel_path := filepath.Join(parent, fname)
		abs_path := filepath.Join(root, rel_path)

		log.Printf("%s resolves to %s\n", url.Path, abs_path)

		fh, err := os.Open(abs_path)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Access-Control-Allow-Origin", "*")
		rsp.Header().Set("Content-Type", "application/json")

		io.Copy(rsp, fh)
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
