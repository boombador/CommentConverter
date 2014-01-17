
package formats

import (
    "encoding/xml"
    "fmt"
    "strings"
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
    Forum string `xml:"forum"`
    Title string `xml:"title"`
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
    Message string `xml:"message"`
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
    Parent DsqIDNode `xml:"parent"`
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
