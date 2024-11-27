package jstore

import (
	"encoding/json"
	"goStudy/pkg/localStore/util"
	"log"
	"os"
)

type Storage struct {
	Filename    string             `json:"filename"`
	ParsedLinks []string           `json:"parsedLinks"`
	DataLinks   map[string][]int64 `json:"data"`
}

func New(filename string) *Storage {
	s := &Storage{Filename: filename}
	s.Load()
	return s
}

func (ls *Storage) Load() error {
	file := util.GetStore(ls.Filename)
	err := json.Unmarshal(file, ls)
	if err != nil {
		log.Println("Error unmarshalling json:", err)
	}
	if ls.DataLinks == nil {
		ls.DataLinks = make(map[string][]int64)
	}
	if ls.ParsedLinks == nil {
		ls.ParsedLinks = make([]string, 0)
	}
	return nil
}

func (ls *Storage) Save() error {
	j, err := json.Marshal(ls)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(ls.Filename, j, 0777)
	if err != nil {
		return err
	}
	return nil
}
func (ls *Storage) Links() []string {
	return ls.ParsedLinks
}
func (ls *Storage) Data() map[string][]int64 {
	return ls.DataLinks
}
func (ls *Storage) SetLinks(links []string) {
	ls.ParsedLinks = links
}
func (ls *Storage) SetData(data map[string][]int64) {
	ls.DataLinks = data
}
