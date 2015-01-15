package main

import (
    "fmt"
    "net/http"
    "os"
	"io/ioutil"
	"log"
)

type Page struct {
    Title string
    Body  []byte
}

func main() {
    http.HandleFunc("/hello", hello)
	http.HandleFunc("/", bluegreen)
    log.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
	  log.Println("ERROR serving the application")
      panic(err)
    }
}

func loadPage() (*Page, error) {
    filename := "blue-green.html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
	log.Println("ERROR reading the file")
        return nil, err
    }
    return &Page{Body: body}, nil
}

func bluegreen(res http.ResponseWriter, req *http.Request) {
	page, _ := loadPage()
	fmt.Fprintf(res, "%s", page.Body)
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}