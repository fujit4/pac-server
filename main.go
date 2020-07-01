package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	var port = flag.String("port", "8282", "port number")
	var rootdir = flag.String("root", ".", "root directory of fileserver")
	// Simple static webserver:
	http.Handle("/", http.FileServer(http.Dir(*rootdir)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
