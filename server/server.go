package main

import (
	"log"
	"net/http"
	"io"
)


func ajax(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "New status dude.")
}
func main() {

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(""))))
	http.HandleFunc("/ajax",ajax)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "root.html")
	})
	log.Println("Listening...")
	http.ListenAndServe(":69", nil)
}
