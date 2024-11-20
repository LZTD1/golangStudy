package reverseIndex

import (
	"goStudy/pkg/crawler"
	"testing"
)

func TestIndex_Add(t *testing.T) {
	tests := []struct {
		name     string
		i        ReverseIndex
		arg      crawler.Document
		expected ReverseIndex
	}{
		{
			name: "Test1 - Single Document",
			i:    make(ReverseIndex),
			arg: crawler.Document{
				ID:    1,
				URL:   "http://google.com",
				Title: "Search everything",
			},
			expected: ReverseIndex{
				"search":     {1},
				"everything": {1},
			},
		},
		{
			name: "Test2 - Adding Same Document",
			i: ReverseIndex{
				"search":     {1},
				"everything": {1},
			},
			arg: crawler.Document{
				ID:    1,
				URL:   "http://google.com",
				Title: "Search everything",
			},
			expected: ReverseIndex{
				"search":     {1},
				"everything": {1},
			},
		},
		{
			name: "Test3 - Same Documents with Same Word",
			i:    make(ReverseIndex),
			arg: crawler.Document{
				ID:    1,
				URL:   "http://test.com",
				Title: "Search for for everything",
			},
			expected: ReverseIndex{
				"search":     {1},
				"for":        {1},
				"everything": {1},
			},
		},
		{
			name: "Test4 - Different Documents with Same Word",
			i: ReverseIndex{
				"search":     {1},
				"everything": {1},
			},
			arg: crawler.Document{
				ID:    2,
				URL:   "http://test.com",
				Title: "Search for for everything",
			},
			expected: ReverseIndex{
				"search":     {1, 2},
				"for":        {2},
				"everything": {1, 2},
			},
		},
		{
			name: "Test5 - Empty Title",
			i:    make(ReverseIndex),
			arg: crawler.Document{
				URL:   "http://empty.com",
				Title: "",
			},
			expected: ReverseIndex{},
		},
		{
			name:     "Test6 - Adding Empty Document",
			i:        make(ReverseIndex),
			arg:      crawler.Document{},
			expected: ReverseIndex{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.Add(tt.arg)

			for word, index := range tt.expected {
				if got := tt.i[word]; !equalDocuments(got, index) {
					t.Errorf("ReverseIndex[%q] = %v, want %v", word, got, index)
				}
			}
		})
	}
}
func equalDocuments(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
