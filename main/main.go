package main

import (
	"fmt"
	"net/http"
)

// type Handler interface{
// ServeHTTP (ResponseWriter, *Request)

type MyWebserverType bool

func (m MyWebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `SUKAAA!`)
}

func main() {
	var k MyWebserverType
	http.ListenAndServe("localhost:8080", k)

}
