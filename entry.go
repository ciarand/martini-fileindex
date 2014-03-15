package fileindex

import (
	"github.com/dustin/go-humanize"
	"net/url"
	"os"
)

type Entry struct {
	Info os.FileInfo
	Url  *url.URL
	Name string
}

func NewEntry(info os.FileInfo) (*Entry, error) {
	name := info.Name()
	url, err := url.Parse(name)
	if err != nil {
		return nil, err
	}

	return &Entry{
		Info: info,
		Url:  url,
		Name: name,
	}, nil
}

func (e *Entry) HumanSize() string {
	return humanize.Bytes(uint64(e.Info.Size()))
}

func (e *Entry) HumanTime() string {
	return humanize.Time(e.Info.ModTime())
}

func (e *Entry) IsDir() bool {
	return e.Info.IsDir()
}
