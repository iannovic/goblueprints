package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//ServeHTTP handles the HTTP Request.
//templateHandler satisfies the http Handler interface with its implementation of ServeHTTP, thus can be used in http.Handle
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("current working directory: ", dir)
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}
func main() {
	//root
	http.Handle("/", &templateHandler{filename: "chat.html"})

	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
