package main

import (
	"flag"
	"fmt"
	"goStudy/pkg/crawler"
	"goStudy/pkg/crawler/spider"
	"strings"
)

func main() {
	f := flag.String("s", "", "Contains word")
	d := flag.Int("d", 2, "Depth scanning")
	flag.Parse()
	if *f != "" {
		fmt.Printf("Установлен флаг - %s\n", *f)
	}

	crawl := spider.New()

	ch := make(chan []crawler.Document)

	go func(ch chan []crawler.Document, d int) {
		fmt.Println("Запуск crawler`а по https://go.dev")
		r, _ := crawl.Scan("https://go.dev", d)
		ch <- r
	}(ch, *d)
	go func(ch chan []crawler.Document, d int) {
		fmt.Println("Запуск crawler`а по https://devdocs.io/go")
		r, _ := crawl.Scan("https://devdocs.io/go/", d)
		ch <- r
	}(ch, *d)

	r := []crawler.Document{}
	for i := 0; i < 2; i++ {
		select {
		case docs := <-ch:
			for _, doc := range docs {
				if strings.Contains(strings.ToLower(doc.Title), *f) {
					r = append(r, doc)
				}
			}
		}
	}

	fmt.Printf("Собрано %d ссылок", len(r))
}
