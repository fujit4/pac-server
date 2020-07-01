package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var port = flag.String("port", "8282", "port number")
	var rootdir = flag.String("root", ".", "root directory of fileserver")
	fmt.Println("start" + " " + "port:" + *port + " " + "root dir:" + *rootdir)
	// Simple static webserver:
	http.Handle("/", http.FileServer(http.Dir(*rootdir)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
