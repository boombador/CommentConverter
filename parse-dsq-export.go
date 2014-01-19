package main

import (
    "encoding/xml"
    "fmt"
    dsq "github.com/boombador/disqus/formats"
    "io/ioutil"
    "os"
)

func main() {
    xmlFile, err := os.Open("xml/nextjump-disqus-comments.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)
    var q dsq.Disqus
    xml.Unmarshal(b, &q)

    // map of thread id to slice of post indices
    var newThreads []dsq.Thread
    var threadMap = make(map[int][]int)
    for ip, post := range q.PostList {
        id := post.Thread.Id
        threadMap[id] = append(threadMap[id], ip)
    }
    for _, thread := range q.ThreadList {
        id := thread.Id
        _, hasPost := threadMap[id]
        if hasPost {
            newThreads = append(newThreads, thread)
        }
    }

    out := dsq.Disqus{xml.Name{"", "disqus"}, "http://disqus.com", q.CategoryList, newThreads, q.PostList}
    newXML, _ := xml.MarshalIndent(&out, "", "    ")
    fmt.Println(`<?xml version="1.0" encoding="utf-8"?>`)
    fmt.Printf("%s", newXML)
}
