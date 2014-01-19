
package formats

import (
    "encoding/xml"
    "fmt"
)

type Rss struct {
    XMLName xml.Name `xml:"rss"`
    // Version []string `xml:"version,attr"`
    Channel Channel `xml:"channel"`
}

type Channel struct {
    // XMLName xml.Name `xml:"channel"`
    Items []Item `xml:"item"`
}

type Item struct {
    XMLName xml.Name `xml:"item"`
    Title string `xml:"title"`
    Link string `xml:"link"`
    Content string `xml:"content:encoded"`
    ThreadIdentifier string `xml:"dsq:thread_identifier"`
    PostDataGmt string `xml:"dsq:post_data_gmt"`
    CommentStatus string `xml:"wp:comment_status"`
    Comments []Comment `xml:"wp:comment"`
}

type Comment struct {
    XMLName xml.Name `xml:"wp:comment"`
    Id int `xml:"wp:comment_id"`
    Author string `xml:"wp:comment_author"`
    Email string `xml:"wp:comment_author_email"`
    AuthorURL string `xml:"wp:comment_author_url"`
    IpAddress string `xml:"wp:comment_author_IP"`
    Datetime string `xml:"wp:comment_date_gmt"`
    Content string `xml:"wp:comment_content"`
    Approved int `xml:"wp:comment_approved"`
    Parent int `xml:"wp:comment_parent"`
}

func (i Item) String() string {
    return fmt.Sprintf("[%s] %s", i.Link, i.Title)
}
func (c Comment) String() string {
    return fmt.Sprintf("%s\nby - %s <%s>; site=\"%s\"", c.Content, c.Author, c.Email, c.AuthorURL)
}
