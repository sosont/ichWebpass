package main

import (
	"io"
	"net/http"
	//"strings"
	"time"
)

var (
	server = &http.Server{
		Addr:           ":9090",
		Handler:        &ppserver{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	handlersMap = make(map[string]HandlersFunc)
)

type ppserver struct {
}

func (*ppserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := handlersMap[r.URL.String()]; ok {
		h(w, r)
	}
	//io.WriteString(w, "URL"+r.URL.String())
}

func f1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "111111111111")
}

func f2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "2222222222222")
}

type HandlersFunc func(http.ResponseWriter, *http.Request)

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {

	handlersMap["/hello"] = Hello
	handlersMap["/f1"] = f1
	handlersMap["/f2"] = f2

	server.ListenAndServe()

}
