package moviesddl

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	config, err := NewConfig("testdata/config/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#+v", config)
}
