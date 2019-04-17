package main

import (
    "fmt"
    "net/http"
)

const (
    SERVER_DOMAIN = "localhost"
    SERVER_PORT = 8080
    RESPONSE_TEMPLATE = "hello"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Constent-Type", "text/html")
    w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
    w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
    http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
    http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
}
