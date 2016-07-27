package main

import (
	"fmt"
	"net/http"
        "net/http/httputil"
        "os"
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
    dump, err := httputil.DumpRequest(request, true)
    if err != nil {
	http.Error(responseWriter, fmt.Sprint(err), http.StatusInternalServerError)
	return
    }    
    fmt.Fprintf("%q",dump)
    fmt.Fprintf(responseWriter, "OK")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
