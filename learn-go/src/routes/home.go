package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"text/template"

	"learn-go/src/db"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// getting Current Working Directory
	rootDir, err := os.Getwd()
	if err != nil {
		// responds with an HTTP 500 Internal Server Error
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	// constructing te template file path into a single file path string
	filename := path.Join(rootDir, "templates", "home.html")
	// parsing the template file
	// template.Must handles parsing errors
	tmpl := template.Must(template.ParseFiles(filename))
	// writes  the output directly to 'w'
	tmpl.Execute(w, nil)
	// the template rendered output is sent as HTTP response
}

func Table(w http.ResponseWriter, r *http.Request) {
	df := db.GetPlayers()
	// send a json represenatation of df. which is slices of struct Table
	// structure that can be encoded as JSON; slice of mas or a custom struct
	// 'ResponseWriter' can be used to send headers, write the response body and set status code
	// Header() nethod returns a map of header keys to their values(metadata)
	// .Set>> sets header filed and value is json-formatted data
	w.Header().Set("content-type", "application/json")
	// 'df' is serialized to JSON and written to the response body.
	json.NewEncoder(w).Encode(df)
}

// Common Content-Type values:

// application/json: For JSON data.
// text/html: For HTML content.
// text/plain: For plain text.
// application/xml: For XML data.
// multipart/form-data: For form submissions that include files.
