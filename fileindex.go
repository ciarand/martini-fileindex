package fileindex

import (
	"errors"
	"fmt"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"net/url"
	"os"
)

func ListFiles(dirpath string) (handler martini.Handler) {
	dir := http.Dir(dirpath)

	return func(w http.ResponseWriter, r *http.Request, log *log.Logger) {
		uri := r.URL.Path

		handle, err := dir.Open(uri)
		if err != nil {
			return
		}
		defer handle.Close()

		err = needsIndex(handle)
		if err != nil {
			return
		}

		log.Println("[Fileindex] Serving file index from " + uri)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<pre>\n")
		printFiles(w, handle)
		fmt.Fprintf(w, "</pre>\n")
	}
}

// Prints all the files (with a linebreak) to the provided writer
func printFiles(w http.ResponseWriter, handle http.File) {
	// While we have files
	for {
		dirs, err := handle.Readdir(100)

		if err != nil || len(dirs) == 0 {
			return
		}

		// For each one of them, print a link
		for _, d := range dirs {
			str, err := printOneFile(d)

			if err != nil {
				break
			}

			fmt.Fprintf(w, str)
		}
	}
}

// Prints one file link, complete with URL escaping
func printOneFile(d os.FileInfo) (str string, err error) {
	name := d.Name()
	if d.IsDir() {
		name += "/"
	}

	// escape the URL
	url, err := url.Parse(name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<a href=\"%s\">%s</a>\n", url, name), nil
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
