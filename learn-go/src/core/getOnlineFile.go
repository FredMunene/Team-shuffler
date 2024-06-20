package core

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func FetchContent(url string, filename string) error {
	// create the file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	// create a client
	// handles redirects,
	// current url is made to match with Path(redirected url changes maybe)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
		// a 30-second timeout
		Timeout: 30 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}
