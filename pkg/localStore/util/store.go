package util

import (
	"os"
)

func GetStore(filename string) []byte {
	f, err := os.ReadFile(filename)
	if err != nil {
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		data, _ := os.ReadFile(f.Name())
		f.Close()
		return data
	}
	return f
}
func WriteStore(data []byte, filename string) {
	err := os.WriteFile(filename, data, 0777)
	if err != nil {
		panic(err)
		return
	}
}
