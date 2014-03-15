package fileindex

import (
	"errors"
	"github.com/codegangsta/martini"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func ListFiles(dirpath string) (handler martini.Handler) {
	dir := http.Dir(canonicalDirPath(dirpath))

	return func(w http.ResponseWriter, r *http.Request, log *log.Logger) {
		uri := r.URL.Path

		handle, err := dir.Open(uri)
		if err != nil {
			log.Println(err)
			return
		}
		defer handle.Close()

		log.Println("[Fileindex] Serving directory index of", dir, "at", uri)

		writeHtml(w, handle, log)
	}
}

func writeHtml(w http.ResponseWriter, handle http.File, log *log.Logger) {
	err := needsIndex(handle)
	if err != nil {
		return
	}

	dirinfo, err := handle.Stat()
	if err != nil {
		return
	}

	err = executeTemplate(w, dirinfo)
	if err != nil {
		log.Println(err)
		return
	}
}

func executeTemplate(w http.ResponseWriter, dirinfo os.FileInfo) error {
	infos, err := ioutil.ReadDir(dirinfo.Name())
	if err != nil {
		infos = make([]os.FileInfo, 0)
	}

	entries := make([]*Entry, len(infos))

	for key, info := range infos {
		entry, err := NewEntry(info)
		if err != nil {
			continue
		}

		entries[key] = entry
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	data, err := Asset("templates/page.html")
	if err != nil {
		return err
	}

	t, err := template.New("document").Parse(string(data))
	if err != nil {
		return err
	}

	params := struct {
		Entries []*Entry
		Title   string
	}{
		entries,
		dirinfo.Name(),
	}

	return t.Execute(w, params)
}

// Checks the provided URI and file handle to see if we actually need to
// provide an index.
func needsIndex(handle http.File) (err error) {
	stat, err := handle.Stat()
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return errors.New("Not a directory")
	}

	return nil
}

func canonicalDirPath(dirpath string) string {
	if !strings.HasSuffix(dirpath, "/") {
		dirpath += "/"
	}

	return dirpath
}
