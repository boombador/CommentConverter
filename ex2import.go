package main

import (
    "encoding/xml"
    "fmt"
    "strings"
    "io/ioutil"
    "os"
)

type Disqus struct {
    XMLName xml.Name `xml:"disqus"`
    Xmlns string `xml:"http://disqus.com/disqus-internals xmlns,attr"`
    // Xmlns string `xml:"http://disqus.com/disqus-internals xmlns,attr"`
    CategoryList []Category `xml:"category"`
    ThreadList []Thread `xml:"thread"`
    PostList []Post `xml:"post"`
}

type Category struct {
    Id int `xml:"http://disqus.com/disqus-internals id,attr"`
    Title string `xml:"title"`
    Forum string `xml:"forum"`
    Default bool `xml:"isDefault"`
}
type Author struct {
    Email string `xml:"email"`
    Name string `xml:"name"`
    Anonymous bool `xml:"isAnonymous"`
    Username string `xml:"username"`
}
type Thread struct {
    Id int `xml:"http://disqus.com/disqus-internals id,attr"`
    Forum string `xml:"forum"`
    Category DsqIDNode `xml:"category"`
    Link string `xml:"link"`
    Title string `xml:"title"`
    CreatedAt string `xml:"createdAt"`
    Author Author `xml:"author"`
    IPAddress string `xml:"ipAddress"`
    Closed bool `xml:"isClosed"`
    Deleted bool `xml:"isDeleted"`
}
type Post struct {
    Id int `xml:"http://disqus.com/disqus-internals id,attr"`
    Message string `xml:"message"`
    CreatedAt string `xml:"createdAt"`
    Deleted bool `xml:"isDeleted"`
    Spam bool `xml:"isSpam"`
    Author Author `xml:"author"`
    IPAddress string `xml:"ipAddress"`
    Thread DsqIDNode `xml:"thread"`
}
type DsqIDNode struct {
    XMLName xml.Name
    Id int `xml:"http://disqus.com/disqus-internals id,attr"`
}

func (a Author) String() string {
    return fmt.Sprintf("[%s]\t%s - %s (anon: %t)", a.Username, a.Name, a.Email, a.Anonymous)
}
func (c Category) String() string {
    return fmt.Sprintf("[%d]\t%s\t\t- from '%s' (default: %t)",
        c.Id, c.Title, c.Forum, c.Default)
}
func (t Thread) String() string {
    return fmt.Sprintf("[%d]\t%s - %s (closed: %t, deleted: %t)", t.Id, t.Title, t.Link, t.Closed, t.Deleted)
}
func (p Post) String() string {
    return fmt.Sprintf("[%d] threadID: %d by %s\n%s",
        p.Id, p.Thread.Id, p.Author, strings.Trim(p.Message, " \n\t"))
}

func main() {
    xmlFile, err := os.Open("nextjump-disqus-comments.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)

    var q Disqus
    xml.Unmarshal(b, &q)

    newXML, _ := xml.MarshalIndent(&q, "", "  ")
    fmt.Printf("%s", newXML)

    /*
    type Person struct {
        Name  string
        Likes []string
    }
    var people []*Person

    likes := make(map[string][]*Person)
    for _, p := range people {
        for _, l := range p.Likes {
            likes[l] = append(likes[l], p)
        }
    }
    */

    // var threadMap = make(map[Thread][]Post)
    // for _, post := range q.PostList {
        // id := post.Thread.Id
        // threadMap[id] = nil //post.

        // add to mapping of thread => posts

        // fmt.Printf("%s\n", post)
    // }
    // iterate through threads
    //     only keep those with entries
/*
    fmt.Println("Categories")
    for _, category := range q.CategoryList {
        fmt.Printf("%s\n", category)
    }
    fmt.Println("Threads")
    for _, thread := range q.ThreadList {
        fmt.Printf("%s\n", thread)
    }
    fmt.Println("Posts")
    for _, post := range q.PostList {
        fmt.Printf("%s\n", post)
    }
*/
}
