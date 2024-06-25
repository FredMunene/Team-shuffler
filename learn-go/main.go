package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"learn-go/src/core"
	handler "learn-go/src/routes"
)

const (
	PORT = 8765
	HOST = "localhost"
	APP  = "Team Management App"
)

func main() {
	// create a Players file with all players
	err := core.FetchContent()
	if err != nil {
		log.Fatal(err)
	}

	// create a TEAMS file
	result, success := core.CreateFile("teams.csv")
	if !success {
		log.Fatalf(result)
	}

	// specifing Host and Port setting (manual)
	port := PORT
	host := HOST
	app := APP

	//
	config, success := core.ReadConfig()
	if !success {
		fmt.Printf("\t\t[[Unable to read configuration file. Using default port: %d]]\n\n", PORT)
	} else {
		// sets default values to retrieved values from configuration file
		port = config.HostPort
		host = config.HostName
		app = config.AppName
	}

	url := fmt.Sprintf("%s:%d", host, port)

	mutex := &sync.Mutex{}

	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/", handler.Home)

	// serve static files
	// http.Dir creates an object to represent the static directory on fs
	// http.FileServer takes Dir object and returns a 'http.Handler' that serves files from directory 'dir'
	// http.StripPrefix strips the prefix from the request URL i.e "static/file.txt" translates to "file.txt"
	// still inside "static" dir
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// routes
	http.HandleFunc("/rules", handler.RulesHandler)
	http.HandleFunc("/table", handler.Table)
	http.HandleFunc("/playerlist", handler.PlayerlistHandler)
	http.HandleFunc("/players", handler.Players)
	http.HandleFunc("/fixtures", handler.Fixtures)
	http.HandleFunc("/shuffle", handler.Shuffle(mutex))
	http.HandleFunc("/register", handler.Register(mutex))
	http.HandleFunc("/toroot", handler.Index)

	fmt.Printf("\n\n\t---[%s]---\n\n\tServer running at %s:%d\n\n", app, host, port)
	http.ListenAndServe(url, nil)
}
