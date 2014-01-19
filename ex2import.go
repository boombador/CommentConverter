package main

import (
    "encoding/xml"
    "fmt"
    dsq "github.com/boombador/disqus/formats"
    "flag"
    "io/ioutil"
    "strings"
    "os"
    "io"
    "encoding/csv"
)

func main() {
    file := flag.String("file", "xml/nextjump-disqus-comments.xml", "Disqus export file path to be converted to import format")
    linkPath := flag.String("map", "", "csv file mapping old links to new")
    flag.Parse()
    xmlFile, err := os.Open(*file)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)
    var q dsq.Disqus
    xml.Unmarshal(b, &q)

    var linkMap = make(map[string]string)
    if *linkPath != "" {
        linkFile, err := os.Open(*linkPath)
        defer linkFile.Close()
        if err != nil {
            panic(err)
        }
        csvReader := csv.NewReader(linkFile)
        for {
            fields, err := csvReader.Read()
            if err == io.EOF {
                break
            } else if err != nil {
                panic(err)
            }
            linkMap[fields[0]] = fields[1]
        }
    }

    var threadMap = make(map[int][]int) // map of thread id to slice of post indices
    for ip, post := range q.PostList {
        id := post.Thread.Id
        threadMap[id] = append(threadMap[id], ip)
    }

    var newThreads []dsq.Thread
    for _, thread := range q.ThreadList {
        id := thread.Id
        _, hasPost := threadMap[id]
        if hasPost {
            newThreads = append(newThreads, thread)
        }
    }

    var newItems []dsq.Item
    for _, t := range newThreads {
        var status string
        if t.Closed {
            status = "closed"
        } else {
            status = "open"
        }
        postIndices := threadMap[t.Id]
        var tComments []dsq.Comment
        for _, i := range postIndices {
            post := q.PostList[i]
            var approved int
            if post.Deleted {
                approved = 0
            } else {
                approved = 1
            }
            tComments = append(tComments,
                dsq.Comment{xml.Name{"", "comment"}, post.Id, post.Author.Name, post.Author.Email, "", post.IPAddress, post.CreatedAt, post.Message, approved, post.Parent.Id})
        }
        mappedLink, mapped := linkMap[t.Link]
        var link string
        if !mapped {
            link = t.Link
        } else {
            link = mappedLink
        }
        newItems = append(newItems, dsq.Item{xml.Name{"", "item"}, t.Title, link, t.Message, "", t.CreatedAt, status, tComments})
    }
    channel := dsq.Channel{newItems}

    out := dsq.Rss{xml.Name{"", "rss"}, channel}
    newXML, _ := xml.MarshalIndent(&out, "", "    ")
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
}
