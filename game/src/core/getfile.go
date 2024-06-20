package core

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile() (err error) {
	// source of our data
	url := "https://bravian1.github.io/players/players.txt"

	os.Mkdir("./storage", os.ModePerm)
	// Create the file
	file, err := os.Create("./storage/players.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	// handling redirects
	// create a client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// change URL of new request(req) 
			// 'Opaque' is set to same value as 'Path' 
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	// Put contents on file
	// fetching the data or webpage at provided url
	response, err := client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	// copy the data from rsponse.Body to file. (online to our system file)
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
