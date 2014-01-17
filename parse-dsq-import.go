
package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
    dsq "github.com/boombador/disqus/formats"
)

func main() {
    xmlFile, err := os.Open("dsq-import.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)
    var q Rss
    xml.Unmarshal(b, &q)

    channel := q.Channel
    for _, item := range channel.Items {
        fmt.Println(item)
        for _, comment := range item.Comments {
            fmt.Println(comment)
        }
    }
}
