package main

import (
	"net/http"
	"flag"
	"io/ioutil"
	"runtime"
	"strconv"
)

type Service struct {
	Api string
}

func (s Service) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c, _ := strconv.Atoi(r.URL.Query().Get("c"))

	ch := make(chan string, c)

	for i := 0; i < c; i++ {
		go func() {
			resp, err := http.Get(s.Api)
			if err != nil {
				ch <- string(err.Error())
				return
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				ch <- string(err.Error())
				return
			}

			ch <- string(body)
		}()
	}

	result := ""
	for i := 0; i < c; i++ {
		result += " " + <-ch
	}

	rw.Write([]byte(result))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var api string
	flag.StringVar(&api, "api", "", "The api endpoint")

	flag.Parse()

	http.Handle("/", Service{Api: api})

	http.ListenAndServe(":80", nil)
}
