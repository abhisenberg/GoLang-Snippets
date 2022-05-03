package chapter1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
Create the simplest of a webserver, which listens to URLs on a specified port number
and print the URL endpoint.
*/
func WebServer1() {
	http.HandleFunc("/", handler1) //Add a generic handler, this one will listen to all the URLs

	result := http.ListenAndServe("localhost:8000", nil) //Start the server on post 8000
	log.Fatal(result)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}

/*
Create a web server that counts the number of APIs hit on it, and displays the count when
the url /count is called.
*/
var mu sync.Mutex
var count int

func WebServer2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/showCount", ShowCount)
	http.ListenAndServe("localhost:8000", nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}

func ShowCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "URL count = %d\n", count)
	mu.Unlock()
}

/*
Create a webserver that prints the details of the request on screen
*/

func WebServer3() {
	http.HandleFunc("/", handler3)
	http.ListenAndServe("localhost:8000", nil)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for key, value := range r.Header { //Iterate over the request headers
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value) //Print the key-value pair of header in fancy way
	}
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr) //Print the remote address

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for key, value := range r.Form { //Get the form from the request and print it
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}

/*
Create a webserver that shows a lissajous gif upon any URL call
*/
func WebServer4() {
	http.HandleFunc("/", handler4) //Add a generic handler, this one will listen to all the URLs

	result := http.ListenAndServe("localhost:8000", nil) //Start the server on post 8000
	log.Fatal(result)
}

func handler4(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}
