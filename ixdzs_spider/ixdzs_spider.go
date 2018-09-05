package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func getJokes(url string, f string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".content").Each(func(i int, s *goquery.Selection) {
		fmt.Println(url)
		ioutil.WriteFile(f, []byte(s.Text()), 0644)
	})
}

func main() {
	for i := 1; i <= 766; i++ {
		url := "http://read.ixdzs.com/52/52247/p" + strconv.Itoa(i) + ".html"
		filePath := "/Users/xxx/Downloads/新驻京办主任-对手/" + strconv.Itoa(i) + ".txt"
		getJokes(url, filePath)
	}
}
