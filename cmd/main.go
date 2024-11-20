package main

import (
	"flag"
	"fmt"
	"goStudy/pkg/crawler"
	"goStudy/pkg/crawler/spider"
	"goStudy/pkg/reverseIndex"
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

	docSort := []crawler.Document{}
	iDoc := make(reverseIndex.ReverseIndex)

	for i := 0; i < 2; i++ {
		select {
		case docs := <-ch:
			for _, doc := range docs {
				iDoc.Add(doc)
				docSort = append(docSort, doc)
			}
		}
	}

	fmt.Printf("Собрано %d ссылок\n", len(docSort))
	fmt.Println(docSort)
	fmt.Println(iDoc)
	fmt.Println("\n")

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
