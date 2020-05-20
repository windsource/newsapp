package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/windsource/newsapp/feed"
)

const feedDefinitionFile = "data/data.json"

var templ *template.Template
var feeds []*feed.Feed

type RetrievedFeed struct {
	Doc     *feed.RSSDocument
	Index   int
	TheFeed *feed.Feed
	Err     error
}

func init() {
	log.SetOutput(os.Stdout)
}

func newsServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	retrievedFeeds := make([]*RetrievedFeed, len(feeds))

	results := make(chan *RetrievedFeed)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for i, singleFeed := range feeds {
		go func(singleFeed *feed.Feed, index int) {
			defer waitGroup.Done()
			doc, err := feed.Retrieve(singleFeed)
			results <- &RetrievedFeed{doc, index, singleFeed, err}
		}(singleFeed, i)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	for result := range results {
		retrievedFeeds[result.Index] = result
	}

	templ.Execute(w, retrievedFeeds)
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
