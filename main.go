package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
