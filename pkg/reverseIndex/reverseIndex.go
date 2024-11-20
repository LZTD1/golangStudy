package reverseIndex

import (
	"goStudy/pkg/crawler"
	"slices"
	"strings"
)

type ReverseIndex map[string][]int64

func (i *ReverseIndex) Add(d crawler.Document) {
	for _, s := range strings.Split(d.Title, " ") {
		s = strings.ToLower(s)

		if val, ok := (*i)[s]; !ok {
			(*i)[s] = []int64{d.ID}
		} else {
			if !slices.Contains(val, d.ID) {
				(*i)[s] = append(val, d.ID)
			}
		}
	}
}
