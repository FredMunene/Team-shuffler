package routes

import (
	"net/http"
	"os"
	"path"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request){
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal server error!", http.StatusInternalServerError)
	}

	// file path to index.html
	filename := path.Join(rootDir,"templates", "index.html")
	// handle errors while parsing contents
	tmpl := template.Must(template.ParseFiles(filename))
	// the template rendered output is sent as HTTP response
	// renders directly to w
	tmpl.Execute(w,nil)
}