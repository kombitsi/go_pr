package actions

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"go_pr/db"
	"github.com/PuerkitoBio/goquery"
)

func Gophers() {
	URL := "https://4gophers.ru"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	linkAll := doc.Find(".item").Find("h2")
	linkText, _ := linkAll.Find("a").Html()
	link, _ := linkAll.Find("a").Attr("href")
	fmt.Println(linkText, link)
	linkEnd := fmt.Sprintf("%s%s", URL, link)
	linkMD5Sum := md5.Sum([]byte(linkEnd))

	text := fmt.Sprintf(`<b>Gopher</b>: <a href\=\"%s\">%s</a>`, linkEnd, linkText)
	fmt.Println(text)
	db.CheckSiteNewsBot(URL, linkEnd, text, fmt.Sprintf("%x", linkMD5Sum))
}
