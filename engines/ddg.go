package engines

import(
	"fmt"
	_ "net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"example.com/Cosmos-Search/web"
)

func search(query string) {
	var url string = "https://duckduckgo.com?q=" + query

	resp := web.getHtml(url)
	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
        log.Fatal(err)
    }

	title := doc.Find("a").Text()
    fmt.Println(title)
}


