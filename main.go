package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)

var headerRe = regexp.MustCompile(`<div class="small_toplink__GmZhY"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	url := "https://www.thepaper.cn/"
	body, err := Fetch(url)
	if err != nil {
		fmt.Println("fetch err:", err)
		return
	}
	doc, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("htmlquery.Parse failed:%v", err)
	}
	nodes := htmlquery.Find(doc, `//div[@class="small_toplink__GmZhY"]/a/h2`)
	for _, node := range nodes {
		fmt.Println("fetch data:", node.FirstChild.Data)
	}
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %v", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Println("fetch error", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
