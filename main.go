package main

import (
	"bufio"
	"encoding/xml"
	"os"
)

type Outline struct {
	XmlURL string `xml:"xmlUrl,attr"`
	Outlines []*Outline `xml:"outline"`
}

func main() {
	var b struct { Body Outline `xml:"body"` }
	err := xml.NewDecoder(bufio.NewReader(os.Stdin)).Decode(&b)
	if err != nil {
		os.Stderr.WriteString(err.Error()+"\n")
		os.Exit(1)
	}
	walk(&b.Body)
}

func walk(r *Outline) {
	if r.XmlURL != "" {
		os.Stdout.WriteString(r.XmlURL+"\n")
	}
	for _, kid := range r.Outlines {
		walk(kid)
	}
}