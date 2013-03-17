package main

import (
	"bufio"
	"encoding/xml"
	"os"
)

type Result struct {
	XMLName xml.Name `xml:"opml"`
	Outlines []Outline `xml:"body>outline"`
}

type Outline struct {
	XmlURL string `xml:"xmlUrl,attr"`
}

func main() {
	var b Result
	err := xml.NewDecoder(bufio.NewReader(os.Stdin)).Decode(&b)
	if err != nil {
		os.Stderr.WriteString(err.Error()+"\n")
		os.Exit(1)
	}

	for _, o := range b.Outlines {
		os.Stdout.WriteString(o.XmlURL+"\n")
	}
}
