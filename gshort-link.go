package main

import (
	"fmt"
	"github.com/markmoudy/gshort-link/shortener"
	// "math/rand"
	"errors"
	"os"
)

var (
	store            = map[string]shortener.Shortener{}
	ErrLinkKeyExists = errors.New("Error: Short Link already exists")
)

func InsertShortLink(dstore map[string]shortener.Shortener, sl shortener.ShortLink) error {
	if _, ok := dstore[sl.ID]; ok {
		return ErrLinkKeyExists
	}
	dstore[sl.ID] = sl
	return nil
}

func main() {
	url := os.Args[len(os.Args)-1]
	sl := shortener.ShortLink{}
	eUrl, err := shortener.GenerateShortUnique(sl, url)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	dID, err := shortener.DecodeShortLink(sl, eUrl)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Printf("Url: %s\nShort Slug: %s\nDecoded Slug: %d\n", url, eUrl, dID)
	os.Exit(0)
}
