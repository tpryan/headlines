package headlines

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLoadFileEmpty(t *testing.T) {

	errEmpty := fmt.Errorf("open : no such file or directory")

	_, err := loadFile("")

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err.Error() != errEmpty.Error() {
		t.Errorf("Expected error (%s), but got (%s)", errEmpty, err)
	}

}

func TestLoadFileNotJson(t *testing.T) {

	errEmpty := fmt.Errorf("invalid character 'p' looking for beginning of value")

	_, err := loadFile("headlines_test.go")

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err.Error() != errEmpty.Error() {
		t.Errorf("Expected error (%s), but got (%s)", errEmpty, err)
	}

}

func TestLoadCacheEmpty(t *testing.T) {

	errNotEmpty := fmt.Errorf("open /subject.json: no such file or directory")

	err := LoadCache("")

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err.Error() != errNotEmpty.Error() {
		t.Errorf("Expected error (%s), but got (%s)", errNotEmpty, err)
	}

}

func TestNewUnloaded(t *testing.T) {

	_, err := New()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err != ErrNotLoaded {
		t.Errorf("Expected error (%s), but got (%s)", ErrNotLoaded, err)
	}

}

func TestNew(t *testing.T) {

	if err := LoadCache("data"); err != nil {
		fmt.Printf("err: %s\n", err)
	}

	cases := []struct {
		Seed int64
		Out  string
	}{
		{0, "Today, neighborhood hipsters in Las Vegas stumbled on to broken windows in a 20 mile radius."},
		{1, "Today, local officials in Ottowa complete work on a completely avoidable easy-to-avoid problem."},
	}

	for _, c := range cases {
		rand.Seed(c.Seed)
		got, err := New()
		if err != nil {
			t.Errorf("Expected no errors, got %s", err)
		}
		if c.Out != got.Sprint() {
			t.Errorf("For %v Expected %s, got %s", c.Seed, c.Out, got)
		}

	}

}
