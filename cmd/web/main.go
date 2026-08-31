package main

import (
	"log"
	"net/http"
	"path/filepath"
	"flag"
)

func main() {
    // Define a new command-line flag with the name 'addr', a default value of ":4000"
    // and some short help text explaning what the flag controls. The value of the
    // flag will be stored in the addr variable at runtime.
    addr := flag.String("addr", ":4000", "HTTP network address")

    // Importantly, we use the flag.Parse() function to parse the commmand-line flag.
    // This reads in the command-line flag value and assigns it to the addr
    // variable. You need to call this *before* you use the addr variable
    // otherwise it will always contain the default value of ":4000". If any errors are
    // encountered during parsing the application will be terminated.
    flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist/view", gistView)
	mux.HandleFunc("/gist/create", gistCreate)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

type neuteredFileSystem struct {
    fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

    f, err := nfs.fs.Open(path)
    if err != nil {
        return nil, err
    }

    s, err := f.Stat()
    if err != nil {
        return nil, err
    }

    if s.IsDir() {
        index := filepath.Join(path, "index.html")
        if _, err := nfs.fs.Open(index); err != nil {
            closeErr := f.Close()
            if closeErr != nil {
                return nil, closeErr
            }

            return nil, err
        }
    }
    return f, nil
}
