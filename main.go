package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	token "github.com/takayukioda/go-md/internal/token"
)

const (
	EXIT_OK = iota
	EXIT_ERROR
	EXIT_HELP
)

const (
	SymbolHeadline rune = '#'
	SymbolEOL           = '\n'
)

type State int

const (
	MightBeNumberList State = iota
)

type Node struct {
	Level int
	Type  token.Type
	Raw   string
	Value string
}

func IsBullet(t rune) {
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("This program need one file name")
		os.Exit(EXIT_ERROR)
	}

	stat, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatalf("Faced an error on checking file status")
		log.Fatalf("Err: %v", err)
		os.Exit(EXIT_ERROR)
	}

	if !stat.Mode().IsRegular() {
		log.Fatalf("Specified file seems not a regular file")
		os.Exit(EXIT_ERROR)
	}

	raw, err := ioutil.ReadFile(stat.Name())

	var t token.Type = token.Text
	var l int = 0
	var nodes []Node
	var nodeRaw []rune
	var nodeValue []rune

	for _, r := range string(raw) {
		if r == SymbolHeadline {
			fmt.Println("headline symbol", string(r))
			t = token.Headline
			l += 1
			nodeRaw = append(nodeRaw, r)
		}
		if r == SymbolEOL {
			fmt.Println("EoL")
			if t == token.Headline {
				n := Node{
					Level: l,
					Type:  t,
					Raw:   string(nodeRaw),
					Value: string(nodeValue),
				}
				fmt.Printf("%#v\n", n)
				nodes = append(nodes, n)
				t = token.Text
				l = 0
				nodeRaw = nil
				nodeValue = nil
			}
		} else {
			if t == token.Headline && r != SymbolHeadline {
				fmt.Println("headline value")
				nodeRaw = append(nodeRaw, r)
				nodeValue = append(nodeValue, r)
				continue
			}
		}
	}
}
