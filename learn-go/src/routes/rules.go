package routes

import  "net/http"


func RulesHandler( w http.ResponseWriter,  r *http.Request){

	// replying to request with the contents of the named file or directory
	http.ServeFile(w,r,"templates/rules.html")
}

