package util

import (
	"os"
	"testing"
)

const NameFile = "./test.txt"

func Test_getFile_createFile(t *testing.T) {
	got := GetStore(NameFile)
	if len(got) != 0 {
		t.Errorf("getStore(\"./test.txt\")=%v, want empty", got)
	}
	_, err := os.ReadFile(NameFile)
	if err != nil {
		t.Errorf("File don`t created!")
	}

	afterAll()
}

func Test_getFile_getExisted(t *testing.T) {
	GetStore(NameFile)
	os.WriteFile(NameFile, []byte("hello world"), 0777)
	got := GetStore(NameFile)
	if len(got) == 0 {
		t.Errorf("getStore(\"./test.txt\")=%v, dont want empty", got)
	}
	if string(got) != "hello world" {
		t.Errorf("getStore(\"./test.txt\")=%v, dont equals hello world", got)
	}

	afterAll()
}

func TestWriteStore(t *testing.T) {
	GetStore(NameFile)
	WriteStore([]byte("hello world"), NameFile)

	got := GetStore(NameFile)
	if string(got) != "hello world" {
		t.Errorf("getStore(\"./test.txt\")=%v, dont equals hello world", got)
	}
	afterAll()
}

func afterAll() {
	os.Remove(NameFile)
}
