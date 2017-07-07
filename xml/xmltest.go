package main

import "encoding/xml"

import "log"
import "fmt"

func main() {
	type OneMessage struct {
		OneLine string
		Country string
	}
	type Message struct {
		TheMessage []OneMessage `xml:"message"`
	}
	type Document struct {
		XMLName  xml.Name `xml:"document"`
		XMLUrl   string   `xml:"xmlurl,attr"`
		Messages Message  `xml:"messages"`
	}
	worldHellos := map[string]string{
		"Chinese":    "你好世界",
		"Dutch":      "Hallo wereld",
		"English":    "Hello world",
		"French":     "Bonjour monde",
		"German":     "Hallo Welt",
		"Greek":      "γειά σου κόσμος",
		"Italian":    "Ciao mondo",
		"Japanese":   "こんにちは世界",
		"Korean":     "여보세요 세계",
		"Portuguese": "Olá mundo",
		"Russian":    "Здравствулте мир",
		"Spanish":    "Hola mundo",
	}

	h := &Document{
		XMLUrl: "http://www.w3.org/2001/XMLShema-instance",
	}

	for k, v := range worldHellos {
		h.Messages.TheMessage = append(h.Messages.TheMessage, OneMessage{
			OneLine: v,
			Country: k,
		})
	}

	//xmlstring, err := xml.MarshalIndent(h, "", "  ")
	xmlstring, err := xml.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	xmlstring = []byte(xml.Header + string(xmlstring))
	fmt.Printf("%s\n", xmlstring)

}
