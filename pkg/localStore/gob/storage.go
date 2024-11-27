package gob

import (
	"encoding/gob"
	"goStudy/pkg/localStore/util"
	"os"
)

type Store struct {
	Filename    string
	ParsedLinks []string
	DataLinks   map[string][]int64
}

func New(filename string) *Store {
	s := &Store{Filename: filename}
	s.Load()
	return s
}

func (ls *Store) Load() error {
	util.GetStore(ls.Filename)
	file, _ := os.Open(ls.Filename)
	defer file.Close()

	dec := gob.NewDecoder(file)
	return dec.Decode(ls)
}

func (ls *Store) Save() error {
	file, err := os.Create(ls.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	return enc.Encode(ls)
}

func (ls *Store) Links() []string {
	return ls.ParsedLinks
}

func (ls *Store) Data() map[string][]int64 {
	return ls.DataLinks
}

func (ls *Store) SetLinks(links []string) {
	ls.ParsedLinks = links
}
func (ls *Store) SetData(data map[string][]int64) {
	ls.DataLinks = data
}
