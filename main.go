package main

import (
	"github.com/Xuanwo/go-language"
	recordjar "github.com/Xuanwo/go-record-jar"
)

type Meta struct {
	FileDate string `json:"File-Date"`
}

func main() {
	data := recordjar.Parse(downloadFromIANA())

	meta := &Meta{}
	tags := make([]language.Tag, 0)

	meta.FileDate = data[0]["File-Date"][0]
	for _, v := range data[1:] {
		tags = append(tags, parseRecordJar(v))
	}

	writeIntoJSON(meta, tags)
	writeIntoMinifiedJSON(meta, tags)
}
