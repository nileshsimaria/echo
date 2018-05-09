package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	j    = flag.Bool("json", false, "JSON")
	port = flag.Int("port", 5001, "Listen on port")
)

func main() {
	flag.Parse()
	http.ListenAndServe(fmt.Sprintf(":%d", *port),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()

			if *j {
				var buf bytes.Buffer
				if err := json.Indent(&buf, b, " >", "  "); err != nil {
					panic(err)
				}
				log.Println(buf.String())
			} else {
				log.Println(string(b))
			}
		}))
}
