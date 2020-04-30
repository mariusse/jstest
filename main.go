package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("*.html"))
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/js", fs)

	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}