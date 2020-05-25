package main

import (
	"flag"
	"fmt"
	"github.com/couchbase/cbauth"
	"github.com/mattcarabine/health-check/pkg/health"
	"runtime"
	"time"
)

func main() {
	fmt.Println("HELLO!!!")
	port := flag.Int("port",  9900,  "Port to listen on")
	cluster := flag.String("cluster", "127.0.0.1:9000", "Port to talk to cluster manager on")
	flag.Parse()


	fmt.Println(*port)
	go health.NewServer(*port)
	runtime.GOMAXPROCS(1)

	u, p, _ := cbauth.GetHTTPServiceAuth(*cluster)
	health.InitClient(u, p, *cluster)

	fmt.Printf("User: %s, Password: %s\n", u, p)

	for true {
		time.Sleep(5 * time.Second)
	}
}
