package main

import (
	"github.com/chirayu/net/trace"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe("localhost:7070", nil)
	if err != nil {
		log.Fatal("Failed to start")
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tr := trace.New("net-trace", r.URL.Path)
	defer tr.Finish()

	log.Println("Starting")
	defer log.Println("Done")

	tr.LazyPrintf("some event %q happened", "hello")

	select {
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
	case <-time.After(5 * time.Second):
		_, err := io.WriteString(w, "World")
		if err != nil {
			log.Println("Error")
		}
	}

}
