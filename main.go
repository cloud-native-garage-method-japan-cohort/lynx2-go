package main

import (
	"encoding/json"
	"fmt"
	title "lynx2-go/title"
	"net/http"
)

type Response struct {
	Title string `json:"title"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := title.NewTitle()
	res := Response{
		Title: t.Get(),
	}
	responseJSON, _ := json.Marshal(res)

	fmt.Fprintf(w, string(responseJSON))
}
func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {

	http.HandleFunc("/", handlerHello)
	http.HandleFunc("/go/", handler)
	http.ListenAndServe(":8080", nil)
}
