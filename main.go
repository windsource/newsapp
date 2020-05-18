package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/windsource/newsapp/feed"
)

const feedDefinitionFile = "data/data.json"

var templ *template.Template
var feeds []*feed.Feed

type Page struct {
	Title string
	Body  []byte
}

func init() {
	log.SetOutput(os.Stdout)
}

func newsServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	docs := make([]*feed.RSSDocument, 0, 10)

	for _, singleFeed := range feeds {
		doc, _ := feed.Retrieve(singleFeed)
		if doc != nil {
			docs = append(docs, doc)
		}
	}

	templ.Execute(w, docs)
}

func main() {
	log.Println("Starting news app")

	var err error
	templ, err = template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatal(err)
	}

	feeds, err = feed.RetrieveFeeds(feedDefinitionFile)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", newsServer)
	http.ListenAndServe(":8080", nil)
}
