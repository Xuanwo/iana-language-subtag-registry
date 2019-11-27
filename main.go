package main

import (
	"github.com/Xuanwo/go-language"
	recordjar "github.com/Xuanwo/go-record-jar"
)

type File struct {
	FileDate string         `json:"File-Date"`
	Tags     []language.Tag `json:"Tags"`
}

func main() {
	data := recordjar.Parse(downloadFromIANA())

	f := &File{}
	f.FileDate = data[0]["File-Date"][0]
	for _, v := range data[1:] {
		f.Tags = append(f.Tags, parseRecordJar(v))
	}

	writeIntoJSON(f)
}
