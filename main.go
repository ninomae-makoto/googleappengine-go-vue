package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	// Imports the Google Cloud Storage client package.
)

// func init() {
// 	http.HandleFunc("/", root)
// }

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
	// http.Handle("/manifest.json", http.FileServer(http.Dir("dist")))
	// http.Handle("/js", http.FileServer(http.Dir("dist")))
	// http.Handle("/css", http.FileServer(http.Dir("dist")))

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

	// fmt.Printf("test")
	// ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	// projectID := "sealion-list"

	// Creates a client.
	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	// Sets the name for the new bucket.
	// bucketName := "test"

	// Creates a Bucket instance.
	// bucket := client.Bucket(bucketName)
	// fmt.Printf("BucketName %v created.\n", bucket)

	// Creates the new bucket.
	// if err := bucket.Create(ctx, projectID, nil); err != nil {
	// 	log.Fatalf("Failed to create bucket: %v", err)
	// }

	// fmt.Printf("Bucket %v created.\n", bucketName)
}
