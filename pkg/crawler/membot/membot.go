package membot

import (
	"goStudy/pkg/crawler"
)

// Service - имитация служба поискового робота.
type Service struct{}

// New - констрктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int) ([]crawler.Document, error) {

	data := []crawler.Document{
		{
			URL:   "https://yandex.ru",
			Title: "Яндекс",
		},
		{
			URL:   "https://google.ru",
			Title: "Google",
		},
	}

	return data, nil
}
