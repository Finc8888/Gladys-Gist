package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from GladysGist" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Avoid to use subtree pattern for root route
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from GladysGist"))
}

// Add a gistView handler function.
func gistView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a gistCreate handler function.
func gistCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        // If it's not, use the w.WriteHeader() method to send a 405 status
        // code and the w.Write() method to write a "Method Not Allowed"
        // response body. We then return from the function so that the
        // subsequent code is not executed.
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
	w.Write([]byte("Create a new gist"))
}

func main() {
	// Use http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist/view", gistView)
	mux.HandleFunc("/gist/create", gistCreate)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
