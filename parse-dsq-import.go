
package main

import (
    "encoding/xml"
    "fmt"
    "strings"
    "io/ioutil"
    "os"
    dsq "github.com/boombador/disqus/formats"
)

func main() {
    xmlFile, err := os.Open("xml/dsq-import.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)
    var q dsq.Rss
    xml.Unmarshal(b, &q)

    channel := q.Channel
    for _, item := range channel.Items {
        fmt.Println("Item:",item)
        for _, comment := range item.Comments {
            fmt.Println("Comment:",comment)
        }
    }
    newXML, err := xml.MarshalIndent(&q, "", "    ")
    // fmt.Printf("%s", err)

    rssTag := `<rss version="2.0"
      xmlns:content="http://purl.org/rss/1.0/modules/content/"
      xmlns:dsq="http://www.disqus.com/"
      xmlns:dc="http://purl.org/dc/elements/1.1/"
      xmlns:wp="http://wordpress.org/export/1.0/">`
    xmlLen := len(newXML)
    xmlString := string(newXML[:xmlLen])

    fmt.Println(`<?xml version="1.0" encoding="utf-8"?>`)
    xmlString = strings.Replace(xmlString, "<rss>", rssTag, 1)
    fmt.Printf("%s", xmlString)

    // fmt.Printf("%s", newXML)
}
