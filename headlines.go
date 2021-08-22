// Package headlines is a tool for creating randomized headlines.  This package
// uses a set of strings contained in several files in /data, to create a
// randomized headline.
package headlines

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

var (
	cache map[string]RandomList
	types = []string{"subject", "location", "verb", "object"}
)

// ErrNotLoaded is an error returned when there is a problem with the list cache
var ErrNotLoaded = fmt.Errorf("headline cache was not initialized")

func init() {
	fmt.Printf("Random seed generated: \n")
	rand.Seed(time.Now().UnixNano())
}

// RandomList is a custom string slice that allows us to get a random member
// of the list, for the random headline generation.
type RandomList []string

// Get returns a random itme from the list.
func (r RandomList) Get() string {
	min := 0
	max := len(r) - 1
	i := rand.Intn(max-min+1) + min

	return r[i]
}

// Headline contains the content of the headline in various parts of a sentence:
// the Subject, Verb, Object and Location.
type Headline struct {
	Subject  string `json:"subject"`
	Location string `json:"location"`
	Verb     string `json:"verb"`
	Object   string `json:"object"`
}

// Sprint returns a string of the constructed headline.
func (h Headline) Sprint() string {
	return fmt.Sprintf("Today, %s in %s %s %s.", h.Subject, h.Location, h.Verb, h.Object)
}

// New creates and returns a new Headline instance.
func New() (Headline, error) {
	r := Headline{}

	if len(cache) == 0 {
		return r, ErrNotLoaded
	}

	r.Location = cache["location"].Get()
	r.Object = cache["object"].Get()
	r.Subject = cache["subject"].Get()
	r.Verb = cache["verb"].Get()
	return r, nil
}

// LoadCache instructs the package to load up the files containing the lists
// of phrases to randomize.
func LoadCache(dir string) error {
	cache = make(map[string]RandomList)

	for _, v := range types {
		list, err := loadFile(dir + "/" + v + ".json")
		if err != nil {
			return err
		}
		cache[v] = list
	}

	return nil
}

func loadFile(path string) (RandomList, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, err
	}
	var l RandomList
	if err = json.Unmarshal(data, &l); err != nil {
		return []string{}, err
	}

	return l, nil
}
