package storage

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Url struct {
	XMLName xml.Name `xml:"url"`
	Short string `xml:"short"`
	Url string `xml:"url"`
}

type Links struct {
	XMLName xml.Name `xml:"shorter"`
	Urls []Url       `xml:"url"`
}

type XmlStorage struct {
	path string
}

func NewXmlStorage(path string) *XmlStorage {
	return &XmlStorage{path: path}
}

func (x *XmlStorage) Store(short string, url string) error {
	data, err := readFile(x.path)
	if err != nil {
		return errors.New(fmt.Sprint("Opening error, %v", err))
	}

	links, err := parseXml(data)
	if err != nil {
		return err
	}
	links.Urls = append(links.Urls, Url{
		Short: short,
		Url:   url,
	})

	stored, err := xml.Marshal(links)
	if err != nil {
		return errors.New(fmt.Sprint("Xml encode error, %v", err))
	}

	return writeFile(x.path, stored)
}

func (x *XmlStorage) Get(short string) (*string, error) {
	data, err := readFile(x.path)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Opening error, %v", err))
	}

	links, err := parseXml(data)
	if err != nil {
		return nil, err
	}

	for _, url := range links.Urls {
		if url.Short == short {
			return &url.Url, nil
		}
	}

	return nil, nil
}

func parseXml(data []byte) (*Links, error) {
	if len(data) == 0 {
		return &Links{
			Urls: make([]Url, 0),
		}, nil
	}

	var links Links
	if err := xml.Unmarshal(data, &links); err != nil {
		return nil, errors.New(fmt.Sprint("Xml parse error, %v", err))
	}

	return &links, nil
}

func readFile(path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func writeFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0666)
}
