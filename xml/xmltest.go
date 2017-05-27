package main

import (
	"encoding/xml"
	"io"
)
import "os"
import "log"

func main() {
	type OneMessage struct {
		OneLine string
		ID      int
	}
	type Message struct {
		TheMessage []OneMessage `xml:"message"`
	}
	type Hello struct {
		XMLName  xml.Name `xml:"hello"`
		Messages Message  `xml:"messages"`
	}

	h := &Hello{}
	h.Messages.TheMessage = append(h.Messages.TheMessage, OneMessage{OneLine: "bonjour le monde"})
	h.Messages.TheMessage = append(h.Messages.TheMessage, OneMessage{OneLine: "hello world"})
	h.Messages.TheMessage = append(h.Messages.TheMessage, OneMessage{OneLine: "buenas el mundo"})

	filename := "hello.xml"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	xmlWriter := io.Writer(file)
	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	err = enc.Encode(h)
	if err != nil {
		log.Fatal(err)
	}
}
