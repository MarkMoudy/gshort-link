package shortener

import (
	"math/rand"
	"testing"
)

func TestGenerateShortUnique(t *testing.T) {
	su := ShortLink{}
	l, err := GenerateShortUnique(su, "")
	l1, _ := GenerateShortUnique(su, "google1")
	l2, _ := GenerateShortUnique(su, "google1")

	if err == nil {
		t.Fatalf("Expected Error for zero length url, got: %s - %s", l, err.Error())
	}

	if l1 == l2 {
		t.Fatalf("Expected Unique ShortLinkMin but got %s and %s", l1, l2)
	}

	if len(l1) != 6 || len(l2) != 6 {
		t.Fatalf("Expected ShortLinkMin to be 6 chars long, got: %s, %s", l1, l2)
	}
}

func TestDecodeShortLink(t *testing.T) {
	sl := ShortLink{}

	if _, err := DecodeShortLink(sl, ""); err == nil {
		t.Fatalf("Expected Error for zero length token")
	}

	if _, err := DecodeShortLink(sl, "5hgZjV4"); err == nil {
		t.Fatalf("Expected Error for incorrect token length")
	}
}

var chars = []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func genBytesStr(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

var smlUrl = "https://www.foo.com/api/v1/bar/?key=" + genBytesStr(8)
var medUrl = "https://www.foo.com/api/v1/bar/?key=" + genBytesStr(32)
var lrgUrl = "https://www.foo.com/api/v1/bar/?key=" + genBytesStr(256)

func BenchmarkGenerateShortUnique_SmallUrl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateShortUnique(ShortLink{}, smlUrl)
	}
}

func BenchmarkGenerateShortUnique_MedUrl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateShortUnique(ShortLink{}, medUrl)
	}
}

func BenchmarkGenerateShortUnique_LrgUrl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateShortUnique(ShortLink{}, lrgUrl)
	}
}
