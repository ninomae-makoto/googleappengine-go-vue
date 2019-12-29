package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	// Imports the Google Cloud Storage client package.
)

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		if os.Args[1] == "local" {
			// local
			log.Printf("url %s", r.URL.Path[1:])
			http.ServeFile(w, r, "dist"+r.URL.Path)
		} else {
			http.NotFound(w, r)
		}
		return
	}

	var indexTmpl, err = template.ParseFiles("dist/index.html")
	// var indexTmpl, err = template.New("dist/index.html").Delims("[[", "]]").ParseFiles("dist/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = indexTmpl.Execute(w, "test"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// [END func_root]

func main() {

	// log.Printf(os.Args[1])
	http.HandleFunc("/", root)
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]

}
