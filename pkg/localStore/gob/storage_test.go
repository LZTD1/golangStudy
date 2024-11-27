package gob

import (
	"os"
	"testing"
)

const NameFile = "./test.gob"

func Test_getStore(t *testing.T) {
	s := New(NameFile)
	s.Load()

	if s == nil {
		t.Errorf("S is nil")
	}
	if s.Filename != NameFile {
		t.Errorf("Filename is wrong")
	}

	afterAll()
}
func Test_saveStore(t *testing.T) {
	s := New(NameFile)
	s.Load()
	s.ParsedLinks = []string{NameFile, NameFile, NameFile}
	s.Save()

	s2 := New(NameFile)
	s2.Load()

	if len(s2.ParsedLinks) != len(s.ParsedLinks) {
		t.Errorf("ParsedLinks is wrong")
	}

	afterAll()
}

func afterAll() {
	os.Remove(NameFile)
}
