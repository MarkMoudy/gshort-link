package main

import (
	"fmt"
	"github.com/markmoudy/gshort-link/shortener"
	"testing"
)

func TestInsertShortLink(t *testing.T) {
	url := "foo1"
	dStore := map[string]shortener.Shortener{}
	sl := shortener.ShortLink{}

	sl.URL = url
	sl.ShortURL, _ = shortener.GenerateShortUnique(sl, url)
	hashID, _ := shortener.DecodeShortLink(sl, sl.ShortURL)
	sl.ID = fmt.Sprintf("%d", hashID)

	err := InsertShortLink(dStore, sl)
	if err != nil {
		t.Fatalf("Error got: %s, expected: nil", err.Error())
	}

	err = InsertShortLink(dStore, sl)
	if err == nil {
		t.Fatalf("Expected: %s, got: %s", ErrLinkKeyExists, err.Error())
	}
}

func TestInsertShortLink_distribution(t *testing.T) {
	url := "foo2/?="
	dStore := map[string]shortener.Shortener{}
	var collisions []string
	for i := 1; i < 100000; i++ {
		sl := shortener.ShortLink{}
		sl.URL = url
		sl.ShortURL, _ = shortener.GenerateShortUnique(sl, sl.URL)
		hashID, _ := shortener.DecodeShortLink(sl, sl.ShortURL)
		sl.ID = fmt.Sprintf("%d", hashID)

		err := InsertShortLink(dStore, sl)
		if err != nil {

			fmt.Printf("Collision cannot insert(i=%d): %+v, Existing Entry: %+v\n", i, sl, dStore[sl.ID])
			collisions = append(collisions, fmt.Sprintf("%s", "Error cannot insert %+v\nExisting Entry: %+v\ngot: %v, expected: nil", sl, dStore[sl.ID], err))
			// t.Errorf("Error cannot insert %+v\nExisting Entry: %+v\ngot: %v, expected: nil", sl, dStore[sl.ID], err)
		}
	}
	if len(collisions) > 0 {
		t.Fatalf("%d Collisions detected", len(collisions))
	}
}
