package headlines

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var cache map[string]RandomList
var types = []string{"subject", "location", "verb", "object"}

var ErrNotLoaded = fmt.Errorf("headline cache was not initialized")

type Headline struct {
	Subject  string
	Location string
	Verb     string
	Object   string
}

func (h Headline) Sprint() string {
	return fmt.Sprintf("Today, %s in %s %s %s.", h.Subject, h.Location, h.Verb, h.Object)
}

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

type RandomList []string

func (r RandomList) Get() string {
	min := 0
	max := len(r) - 1
	i := rand.Intn(max-min+1) + min

	return r[i]
}

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

func loadFile(path string) ([]string, error) {

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
