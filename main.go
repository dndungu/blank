package main

import "net/http"

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("<html><head></head><body style=\"background-color:#000000\">black</body></html>"))
		},
	)
	http.ListenAndServe(":8080", nil)
}
