package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("LET'S FREAKING GOOOOOOOO!\n"))
	})
	addr := ":80"
	fmt.Println("Example app listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
