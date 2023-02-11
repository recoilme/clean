package main

import (
	"fmt"
	"net/http"

	"github.com/recoilme/clean"
)

func main() {
	head := "<!DOCTYPE html>\n<html>\n<head>\n <meta charset=\"utf-8\"><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/light.min.css\">\n</head>\n<body>"
	body := "<form method=\"GET\">URL: <input type='text' size='30' name='url' value="
	formend := "><input type='submit' value='Go!'></form>"

	footer := "</body></html>"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		content := ""
		link := u.Query().Get("url")

		if link != "" {
			_, content, _ = clean.URI2TXT(link, true) //clean.URI(link, true)
		}
		fmt.Fprintf(w, head+body+link+formend+content+footer)
	})
	http.ListenAndServe(":8080", nil)
}
