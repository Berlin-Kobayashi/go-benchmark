package main

import (
	"net/http"
	"flag"
	"fmt"
	"io/ioutil"
	"runtime"
)

type Service struct {
	Api string
}

func (s Service) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	go func() {
		resp, err := http.Get(s.Api)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
		}

		fmt.Println(string(body))
	}()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var api string
	flag.StringVar(&api, "api", "", "The api endpoint")

	flag.Parse()

	http.Handle("/", Service{Api: api})

	http.ListenAndServe(":80", nil)
}
