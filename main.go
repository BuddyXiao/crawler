package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/buddyxiao/crawler/collect"
	"time"
)

func main() {
	url := "https://www.thepaper.cn/"
	var f collect.Fetcher = collect.BrowserFetch{Timeout: 3 * time.Second}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("read content failed:%v", err)
	}
	doc.Find("div.small_cardcontent__BTALp a h2").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		fmt.Println(i, ":", content)
	})
}
