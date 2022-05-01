package chapter1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Hitting an API endpoint and reading the data using Go's inbuilt net/http module.
Uses ioutil.ReadAll() to read the response body
*/
func FetchURL() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url) //Attempting to get the response
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err:%s", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(response.Body) //If response is found, attempting to read it's body
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Response read err:%s", err)
			os.Exit(1)
		}

		fmt.Printf("%s", body)
	}
}

/*
Uses io.Copy() to copy the response to stdout
*/
func FetchURL2() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url) //Attempting to get the response
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err:%s", err)
			os.Exit(1)
		}

		res, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Response read err:%s", err)
			os.Exit(1)
		}
		fmt.Print(res, response.Status)
	}
}
