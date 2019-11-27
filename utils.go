package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/Xuanwo/go-language"
)

func downloadFromIANA() []byte {
	url := "https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("download http.Get failed for %v", err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("download ioutil.ReadAll failed for %v", err)
	}

	return content
}

func mustGotOne(s []string) string {
	if len(s) != 1 {
		log.Fatal("length of %s is not one")
	}
	return s[0]
}

func parseRecordJar(m map[string][]string) (t language.Tag) {
	// Handle must contain values.
	t.Type = mustGotOne(m["Type"])
	t.Description = m["Description"]
	t.Added = mustGotOne(m["Added"])

	if v, ok := m["Tag"]; ok {
		t.Tag = mustGotOne(v)
	}
	if v, ok := m["Subtag"]; ok {
		t.Subtag = mustGotOne(v)
	}

	if v, ok := m["Deprecated"]; ok {
		t.Deprecated = mustGotOne(v)
	}
	if v, ok := m["Preferred-Value"]; ok {
		t.PreferredValue = mustGotOne(v)
	}
	if v, ok := m["Prefix"]; ok {
		t.Prefix = v
	} else {
		t.Prefix = []string{}
	}
	if v, ok := m["Suppress-Script"]; ok {
		t.SuppressScript = mustGotOne(v)
	}
	if v, ok := m["Macrolanguage"]; ok {
		t.Macrolanguage = mustGotOne(v)
	}
	if v, ok := m["Scope"]; ok {
		t.Scope = mustGotOne(v)
	}
	if v, ok := m["Comments"]; ok {
		t.Comments = mustGotOne(v)
	}

	return
}

func writeIntoJSON(f *File) {
	content, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path.Join("docs", "language-subtag-registry.min.json"), content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	content, err = json.MarshalIndent(f, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path.Join("docs", "language-subtag-registry.json"), content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
