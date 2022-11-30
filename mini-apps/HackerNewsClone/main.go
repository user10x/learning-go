package main

import (
	"HackerNewsClone/hn"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 4000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	log.Println("")

	http.HandleFunc("/", handler(numStories, tpl))
	log.Printf("serving localhost:%d/ handler, top %d stories\n", port, numStories)

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

type storyCache struct {
	numOfStories int
	cache        []item
	expiration   time.Time
	duration     time.Duration
	mutex        sync.Mutex
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {

	sc := storyCache{
		numOfStories: numStories,
		duration:     6 * time.Second,
	}
	// ticker to update the stories after every 3 seconds
	// reciever updates the struct
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			//<-ticker.C moved this down
			temp := storyCache{
				numOfStories: numStories,
				duration:     6 * time.Second,
			}
			temp.getCachedStoriesFaster()
			sc.mutex.Lock()
			sc.cache = temp.cache
			//fmt.Println(temp.cache)
			sc.expiration = temp.expiration
			sc.mutex.Unlock()
			<-ticker.C
		}

	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.getCachedStoriesFaster()
		//log.Println(stories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := templateData{
			Stories:         stories,
			Time:            time.Now().Sub(start),
			NumberOfStories: numStories,
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func getTopStories(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		//http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
		return nil, errors.New("failed to get top stories")
	}

	var stories []item

	at := 0
	for len(stories) < numStories {
		//fmt.Println("loop")
		need := (numStories - len(stories)) * 5 / 4
		stories = append(stories, getStories(ids[at:at+need])...)
		at += need
	}

	return stories[:numStories], nil
}

var (
	cache           []item
	cacheExpiration time.Time
	cacheMutex      sync.Mutex
)

// new version reciever on type storyCache so it can be called on that type
func (sc *storyCache) getCachedStoriesFaster() ([]item, error) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if time.Now().Sub(sc.expiration) < 0 {
		return sc.cache, nil
	}
	stories, err := getTopStories(sc.numOfStories)
	if err != nil {
		return nil, err
	}
	sc.cache = stories
	sc.expiration = time.Now().Add(sc.duration) // set cache timing
	return sc.cache, err
}

// old version -- simple
// keeping for reference
func getCachedStories(numStories int) ([]item, error) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if time.Now().Sub(cacheExpiration) < 0 {
		return cache, nil
	}
	stories, err := getTopStories(numStories)
	if err != nil {
		return nil, err
	}
	cache = stories
	cacheExpiration = time.Now().Add(1 * time.Second) // set cache timing
	return stories, err
}

// adding concurrency
func getStories(ids []int) []item {

	type result struct {
		idx  int
		item item
		err  error
	}
	resultCh := make(chan result)
	var stories []item
	for idx, id := range ids {
		go func(idx, id int) {
			var client hn.Client // avoids race conditions for multiple go routines
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultCh <- result{idx: idx, err: err}
			}
			resultCh <- result{idx: idx, item: parseHNItem(hnItem)}
		}(idx, id)
	}

	var results []result
	for i := 0; i < len(ids); i++ {
		results = append(results, <-resultCh)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})
	for _, res := range results {
		if res.err != nil {
			continue
		}
		if isStoryLink(res.item) {
			stories = append(stories, res.item)
		}
	}
	return stories

}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories         []item
	Time            time.Duration
	NumberOfStories int
}
