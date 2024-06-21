package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/testing", http.HandlerFunc(getHandler))
	http.ListenAndServe(":80", mux)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is your requested resource.")

}
