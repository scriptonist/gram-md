package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/skratchdot/open-golang/open"
)

var filename string

func main() {
	mux := http.NewServeMux()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println(wd)

	box := packr.NewBox("../dist")
	if len(os.Args) < 2 {
		fmt.Println("Usage correctme <filename>")
	}
	filename = filepath.Join(wd, os.Args[1])

	mux.Handle("/", http.FileServer(box))
	mux.HandleFunc("/save", handler)
	// h := cors.AllowAll().Handler(mux)

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("error creating a lister %v", err)
	}
	fmt.Println("Listening on : ", listener.Addr().String())
	fmt.Println("Writing to file : ", filename)
	open.Run("http://" + listener.Addr().String())

	panic(http.Serve(listener, mux))

}

func handler(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	b.ReadFrom(r.Body)
	defer r.Body.Close()

	err := ioutil.WriteFile(filename, b.Bytes(), 0755)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
