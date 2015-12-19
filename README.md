# glink-short - Golang Link Shortener
Simple link shortener written in go.  

## Design
Links come in and short links come out.  This application stores original link, short link, id. Functions to encode, redirect, and store short links. Each short link is unique even if it shares a url. 

Generates a unique 6 unit symbol consisting of a base58 character set to eliminate confusing characters(0,O,I,l).
```
123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
```

## Features(in progress)
[] Stores a Shortened link,
[] keeps statistics about the link, 
[] Shortened links are tied to a user.
[] Redirect functionality. 
[] User Web Interface to add and manage short links, view statistics, etc.

## Usage
```
$ go get -u github.com/markmoudy/gshort-link
$ go run gshort-link.go www.foo.com
```

Test with: (also runs collision checker)
```
go test -v ./...
```

Run benchmarks:
```
go test -v -bench=.
```

