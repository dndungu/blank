package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	colour := struct {
		Code string
		Name string
	}{Code: "#000000", Name: "black"}

	t, err := template.ParseFiles("page.tpl")

	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			t.Execute(w, colour)
		},
	)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BLANKPAGE_PORT")), nil)
}
