package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("started. listening on: " + port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=> NEW REQUEST! <=")
		fmt.Println("referer:", r.Referer())
		fmt.Println("=> <=")

		_, err := w.Write([]byte(`<head><title>malte</title></head><body>moin</body>`))
		if err != nil {
			log.Fatalln(err)
		}
	})
	if err := http.ListenAndServe("localhost:"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
