package main

import (
	"fmt"
	"net/http"
        "net/http/httputil"
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
    dump, err := httputil.DumpRequest(request, true)
    if err != nil {
	http.Error(responseWriter, fmt.Sprint(err), http.StatusInternalServerError)
	return
    }    
    fmt.Printf("%q",dump)
    fmt.Fprintf(responseWriter, "OK")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
