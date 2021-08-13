package headlines

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

var cache map[string]RandomList

var types = []string{"subject", "location", "verb", "object"}

type Headline struct {
	Subject  string
	Location string
	Verb     string
	Object   string
}

type RandomList []string

func (r RandomList) GetValue() string {
	min := 0
	max := len(r) - 1
	i := rand.Intn(max-min+1) + min

	return r[i]
}

func NewHeadline() (Headline, error) {

	r := Headline{}

	r.Location = cache["location"].GetValue()
	r.Object = cache["object"].GetValue()
	r.Subject = cache["subject"].GetValue()
	r.Verb = cache["verb"].GetValue()

	return r, nil
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
	var slice RandomList
	err = json.Unmarshal(data, &slice)
	if err != nil {
		return []string{}, err
	}

	return slice, nil
}
