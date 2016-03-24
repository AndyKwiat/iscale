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
	fs := http.FileServer(http.Dir(""))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/ajax",ajax)
	log.Println("Listening...")
	http.ListenAndServe(":69", nil)
}
