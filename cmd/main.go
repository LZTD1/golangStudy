package main

import (
	"flag"
	"fmt"
	"goStudy/pkg/crawler"
	"goStudy/pkg/crawler/spider"
	"goStudy/pkg/localStore"
	"goStudy/pkg/localStore/gob"
	"goStudy/pkg/reverseIndex"
	"slices"
)

var urls = []string{
	"https://go.dev",
	"https://devdocs.io/go/",
}

func main() {
	var s localStore.Store = gob.New("./cache.gob")

	f := flag.String("s", "", "Contains word")
	d := flag.Int("d", 2, "Depth scanning")
	flag.Parse()
	if *f != "" {
		fmt.Printf("Установлен флаг - %s\n", *f)
	}

	crawl := spider.New()

	links := s.Links()
	urlsToCrawl := []string{}

	for _, url := range urls {
		if !slices.Contains(links, url) {
			urlsToCrawl = append(urlsToCrawl, url)
		}
	}

	ch := make(chan []crawler.Document, len(urlsToCrawl))

	for _, url := range urlsToCrawl {
		go func(ch chan []crawler.Document, d int, url string) {
			fmt.Printf("Запуск crawler`а по %s\n", url)
			r, _ := crawl.Scan(url, d)
			ch <- r
		}(ch, *d, url)
	}

	docSort := []crawler.Document{}
	iDoc := make(reverseIndex.ReverseIndex)

	if len(urlsToCrawl) == 0 {
		iDoc = s.Data()
	} else {
		for range urlsToCrawl {
			select {
			case docs := <-ch:
				for _, doc := range docs {
					iDoc.Add(doc)
					docSort = append(docSort, doc)
				}
			}
		}
		close(ch)
		s.SetData(iDoc)
		s.SetLinks(urlsToCrawl)
	}

	fmt.Printf("Собрано %d ссылок\n", len(docSort))

	fmt.Println()
	fmt.Println(s.Links())
	fmt.Println(s.Data())
	fmt.Println()

	s.Save()
	val, ok := iDoc[*f]
	if *f != "" && ok {
		for _, i2 := range val {
			fmt.Println(docSort[binarySearch(docSort, i2)].Title)
		}
	} else {
		fmt.Println("Такого слова не найдено !")
	}

}

func binarySearch(arr []crawler.Document, target int64) int64 {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid].ID == target {
			return int64(mid)
		}
		if arr[mid].ID > target {
			high = mid - 1
		}
		if arr[mid].ID < target {
			low = mid + 1
		}
	}
	return -1
}
