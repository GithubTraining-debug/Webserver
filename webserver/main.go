package main

import (
	"fmt"
	"net/http"
	"os"
)

func webserver() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":2000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./design.html")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	fmt.Fprintf(
		w,
		string(data),
	)

}

func main() {

	webserver()

}
