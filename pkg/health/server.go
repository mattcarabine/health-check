package health

import (
	"fmt"
	"net/http"
)

func NewServer(port int){
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/summary", SummaryHandler)
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/connections", ConnectionsHandler)
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting http server on %s\n", addr)
	_ = http.ListenAndServe(addr, nil)
}