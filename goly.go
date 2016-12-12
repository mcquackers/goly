package main

import (
	results "github.com/user/goNotifierPkg"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Test struct {
	Name string
}

func serveImage(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, req.URL.Path[1:])
}

func serveCss(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, req, req.URL.Path[1:])
}

func displayTests(page http.ResponseWriter, req *http.Request) {
	page.Header().Set("Content-Type", "text/html")
	page.Header().Set("refresh", "86400")
	t := template.New("dashboard")
	t = template.Must(t.ParseFiles("templates/dashboard.html", "templates/pixel.html", "templates/behavior.html", "templates/segment-list.html", "templates/segment-make.html", "templates/errors.html"))
	t.Execute(page, results.Result)
}

func openAndRead() ([]byte, error) {
	result, err := ioutil.ReadFile("results.dat")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateDisplay(w http.ResponseWriter, req *http.Request) {
	results.UpdateResults()
	http.Redirect(w, req, "/", 302)
}

func main() {
	http.HandleFunc("/", displayTests)
	http.HandleFunc("/img/", serveImage)
	http.HandleFunc("/resources/", serveCss)
	http.HandleFunc("/update", updateDisplay)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
