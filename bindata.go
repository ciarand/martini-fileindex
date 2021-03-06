package fileindex

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"unsafe"
)

func bindata_read(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _templates_page_html = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x53\xcf\x4f\xc2\x30\x14\xbe\xf3\x57\xd4\x9e\xf4\x60\x1b\xe2\x75\xec\x02\x24\x98\xe0\x8f\xc4\x79\xf0\x58\xd6\x37\xd7\xb8\x75\xd8\x16\x13\x6c\xfa\xbf\xdb\xd2\x21\xe0\x18\x18\x13\x7b\x5a\xdf\xfb\xfa\xbd\x1f\xfb\x3e\x6b\x39\x14\x42\x02\xc2\xbc\xc9\x57\x35\x48\x83\x9d\x1b\x24\x17\x93\x87\x71\xf6\xf2\x38\x45\xb3\xec\x6e\x9e\x0e\x92\x12\x18\x4f\x07\xc8\x9f\xc4\x08\x53\x41\x6a\xad\x28\xd0\x25\xbc\x23\x92\x85\x3b\xc2\x04\x5f\x39\x37\x5e\x29\xe5\x29\xac\x85\x4a\x83\x73\xd6\xc6\x6c\xf8\x02\xc9\x9d\x4b\x68\x7c\x1d\x99\x2a\x21\xdf\x90\x82\x6a\x84\xb5\x59\x57\xa0\x4b\x00\x83\x51\xa9\xa0\x18\x61\x4a\x25\x18\x2e\x19\x59\x34\x8d\xd1\x46\xb1\x65\xce\x25\xc9\x9b\x9a\x7e\x07\xe8\x0d\x19\x92\x21\xcd\xb5\xde\xc5\x48\x2d\x3c\x4a\x6b\xec\x7b\xa6\xb1\xe9\x64\xd1\xf0\x75\x5b\x91\x8b\x0f\x94\x57\x4c\xeb\x11\xce\x1b\x69\x98\x1f\x5c\xe1\x98\x8b\xb3\xb1\x85\x9f\xa5\x45\x6c\x2e\x7b\xd9\x88\xd8\x6d\xe2\x30\xae\xba\xc1\xf6\x41\x7a\xcf\x6a\xf0\x93\x97\xfd\x88\x39\xd3\x06\xd5\x0d\x17\x85\x00\x7e\x1a\xfa\x24\x3e\x7b\xc8\x7c\xf4\x47\x13\x01\xd7\x69\x37\x31\xbb\x85\xec\x1f\x6b\x15\x93\xaf\x80\xc8\x54\x1a\x25\x40\x7b\x19\x1c\xef\xa1\x67\xd2\x98\x3c\xb2\x9b\x03\x00\x6b\xff\xaf\x97\xc6\xb3\xaa\x9c\xc3\x5e\x49\x24\x2c\x28\x88\x83\x9d\x60\xa6\xa7\xa8\xcf\xd7\xd5\x4b\x26\x43\xa9\xd9\xaa\x66\x32\x13\xb1\xde\x26\xf8\x5f\x25\x37\x0e\x21\xb7\x7a\x22\x54\xcf\x2a\xbb\x0d\x5e\x9f\xeb\x29\x12\x47\x77\xfd\x92\x73\x3b\x74\x10\xce\xf9\xa1\xdb\x02\xc1\xac\x7f\x58\x4c\x57\x82\x7d\x7c\x1e\x79\x28\x43\x1f\x08\x7e\x6b\x8d\x4a\xbd\x53\x83\x85\x23\x66\x4b\xf0\x15\x00\x00\xff\xff\x68\x18\x70\xc9\xad\x04\x00\x00"

func templates_page_html() ([]byte, error) {
	return bindata_read(
		_templates_page_html,
		"templates/page.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/page.html": templates_page_html,
}
