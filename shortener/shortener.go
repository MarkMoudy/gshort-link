package shortener

import (
	"errors"
	"math"
	"math/rand"
	"strings"
)

var (
	ErrInvalidURLZeroLength   = errors.New("Invalid URL: zero length")
	ErrInvalidTokenZeroLength = errors.New("Invalid Token: zero length")
	ErrInvalidTokenWrongSize  = errors.New("Invalid Token: Incorrect length")
)

type Shortener interface {
	Encode(uint32) string
	Decode(string) uint32
}

func GenerateShortUnique(sl Shortener, url string) (string, error) {
	if url == "" {
		return "", ErrInvalidURLZeroLength
	}

	return sl.Encode(BernsteinHash(url)), nil
}

func DecodeShortLink(sl Shortener, token string) (uint32, error) {
	if token == "" {
		return 0, ErrInvalidTokenZeroLength
	}

	if len(token) > 6 {
		return 0, ErrInvalidTokenWrongSize
	}

	return sl.Decode(token), nil
}

type ShortLink struct {
	ID       string
	URL      string
	ShortURL string
	ShortSlug
}

type ShortSlug struct{}

const base58 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

// BernsteinHash generates a 32-bit hash from a given url.
func BernsteinHash(url string) uint32 {
	h := uint32(5381)
	for _, b := range []byte(url) {
		h = ((h << 6) + h) + uint32(b)
	}

	return h
}

// shuffleAlphabet is a Fisher-Yates shuffle implementation to bring
// some uniqueness to the slugs and prevent collisions.
func shuffleAlphabet(alphabet string) []byte {
	shuffled := []byte(alphabet)
	for i := len(shuffled) - 1; i >= 1; i-- {
		rand := rand.Intn(i)
		shuffled[rand], shuffled[i] = shuffled[i], shuffled[rand]
	}

	return shuffled
}

// Encode takes an id string and returns a token.
// TODO: maybe salt this so it can be recoverable?
func (sl ShortSlug) Encode(id uint32) string {
	token := make([]byte, 0)
	shuffledBase := shuffleAlphabet(base58)
	if id == 0 {
		return string(shuffledBase[0])
	}

	for id > 0 {
		rem := id % uint32(len(shuffledBase))
		token = append(token, shuffledBase[rem])
		id = id / uint32(len(shuffledBase))
	}

	//reverse
	for i := len(token)/2 - 1; i >= 0; i-- {
		opp := len(token) - 1 - i
		token[i], token[opp] = token[opp], token[i]
	}

	return string(token)
}

// Decode takes a shortlink token and returns the original id
func (sl ShortSlug) Decode(token string) uint32 {
	id := int(0)
	pos := float64(len(token)) - 1

	for i := 0; i < len(token); i++ {
		id += strings.Index(string(base58), string(token[i])) * int(math.Pow(float64(len(base58)), pos))
		pos--
	}

	return uint32(id)
}
