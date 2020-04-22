package main

import (
	"flag"
	"net/http"
	"os"
)

// second project, start a server with file system and put code file inside
func main() {
	// dir, _ := os.Getwd()
	// http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
	var dir string
	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to serve") // path is reference
	flag.Parse()

	if *path == "" { // use * to dereference
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}
	// after go run fs.go, can pass arguments "-port=5000 -path=something"
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}

// go run fs.go to run, ctrl + C to stop

// use "go build" cmd to build static-linked, optional: "-o renamed_name"
// then run "go install" and go to go/bin (workplace you compile go code) use "./fs -help" to check

// to build binaray file for linux system:
// use "GOOS=linux GOARCH=arm go build" and just executable on linux
