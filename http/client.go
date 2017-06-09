// Code generated by go-bindata.
// sources:
// client/app.js
// client/index.html
// DO NOT EDIT!

package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clientAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\xcd\x6e\xdb\xb8\x13\xbf\xeb\x29\x06\x6a\x01\xd9\xf8\x3b\x32\x72\x2b\x64\xf8\xd0\x36\xfd\x63\x8b\x45\xe2\x22\xfd\x00\xb6\x41\x0e\x8c\x34\xb6\x89\x48\xa4\x96\xa4\x9c\xa4\x85\xef\x7b\xe9\x43\xf4\xb0\x87\x3d\xec\x2b\xf4\xd2\xe6\xbd\x16\xa4\x28\x59\x1f\x54\xec\xec\x2e\xb0\x3e\x24\x32\xe7\x37\x9c\xaf\x1f\x87\x23\xfb\x85\x44\x90\x4a\xd0\x58\xf9\x33\xcf\xfb\x50\x60\x58\x48\x1c\x7d\x28\xf0\x94\x28\x14\x94\xa4\xe3\x59\x73\xf5\xf5\xf1\x33\x36\xf6\xbc\x98\x33\xa9\x80\x1e\x3f\x63\x30\x07\x86\x37\x60\x45\xa3\xcf\x1e\x00\x40\xca\x63\x92\x62\x04\x81\xbc\x0e\x26\x00\xd3\x29\xbc\x5b\x9c\x2c\x22\x90\xa8\xac\x0c\x96\x82\x67\x10\xa7\x14\x99\x32\x2a\x19\x4a\x49\x56\x28\x23\x28\xb7\xd0\x1f\x64\xcd\x6f\x0d\x54\x77\x59\x7f\x18\xde\xc8\x08\x82\x33\xbc\x91\xc1\xc4\x29\x7d\x87\xb7\xca\x22\x40\xe1\xad\x1a\x80\xbd\x49\x49\x8c\x6b\x9e\x26\x28\x22\x08\xde\xdd\xe5\x68\x96\x81\xb0\x04\x72\x81\x52\x02\x32\x85\xc2\xa1\x2d\x51\x6c\x68\xac\x63\x08\xde\xda\xc7\x61\x94\xf5\xc6\x02\x87\x1c\xb2\x60\x87\x4f\x56\x72\x80\x5b\x6b\x5e\x08\xed\xd3\x22\x47\x46\xd9\xaa\xfc\xee\xc0\x25\xe4\x2e\x82\xe0\x84\xdc\x39\x64\x24\x8b\x20\x78\x7e\xea\x90\xe4\x5a\xf2\xc6\x25\x59\x72\xae\x18\x57\xb8\xcb\x7b\xbd\x34\x14\x6e\x25\x77\xc4\x5b\xab\xee\x0f\x38\x4e\xb9\xc4\x17\x8a\x45\x10\xbc\xd4\x8f\xae\xbc\x92\x0d\x26\x8b\x6b\x5d\x00\xfd\x04\x8b\x9f\x1d\x20\x14\xe2\x8c\x9f\xa3\xcc\x39\x93\x9a\xcd\xaf\xce\xcf\x17\xe7\x11\x9c\x71\x10\x76\x31\x68\xe9\x6c\xeb\x6f\xdb\xdd\x66\xf2\xfa\xf1\x24\x5e\x7c\x62\x24\x73\x95\xa1\x41\x63\xfd\x0f\xb8\x06\x16\x07\x11\xf9\x8c\xe4\xdf\xff\xbc\xff\x5a\xaa\x00\x01\xa9\x52\xf2\xe3\xcb\x41\x5c\x4e\x8b\xfb\x6f\x57\x2e\x77\xda\x54\xfe\xc8\x90\x51\x04\xb9\x0f\xee\xf6\xab\x54\x2b\xf6\x7b\x56\xd3\x59\x24\x94\x91\x1f\x5f\xd8\xf7\x3f\x60\xcd\x13\xca\x5c\x26\x2d\xa7\xf1\xfe\xb7\x21\x52\x9f\x70\x7e\x85\x89\x2b\x87\x25\xb3\x8d\xd8\x45\xa1\x0e\xbd\x4d\x3d\x72\xfe\x89\x7d\xff\x9a\x5d\xbb\x5c\x71\x73\xbb\x0a\xbf\xd2\x3c\x20\x01\x0d\x7a\x7f\x24\x6a\xc3\x05\xbd\xff\xfd\x41\x8a\xbf\x4f\xf9\xfd\x37\xd4\x79\x3a\x88\xe6\x2f\x7f\xfa\xe5\xc5\xf3\xc8\x54\x0b\x05\x30\xe4\x49\xce\x37\x98\x90\x74\x90\xed\x5e\xf9\x77\x3b\x9e\x79\x9e\x37\x9d\xc2\x26\x03\x2a\x61\x53\x20\x50\x26\x15\x61\xb6\x4d\xa1\x10\x5c\x68\x49\xf9\x20\x50\x16\xa9\x02\x81\xaa\x10\x0c\x13\xb8\xba\x03\x72\x4b\xb9\xf4\x96\x05\x8b\x15\xe5\x0c\xe4\x9a\xdf\xbc\xd2\xd8\xd1\x26\x9b\x94\x5a\x63\x7b\x6a\xe8\x12\x46\x66\x21\xac\xce\xe2\xb8\x71\x9e\xf4\xfd\xc4\x53\x0c\x53\xbe\xea\xa0\xc2\x84\x28\x32\x9e\x1d\x82\x94\x8a\xa8\x42\x1e\x86\x5d\x23\x49\x50\x34\xc1\x9b\xcc\x48\x4f\xe5\x0a\xe6\xe0\xf0\xa1\x44\x6e\x01\x53\x89\xad\x68\x7e\x2d\x50\xaa\x7d\xc1\x94\xa0\x01\x6b\x9b\x2c\x7c\xaa\x46\xbe\xed\x33\x61\xab\xbe\xfe\x78\xd6\xb2\xec\x36\x13\x98\xac\x07\x36\xe7\xa1\xdd\xe9\xe1\xe8\x2c\xc8\xee\xee\xb9\x3d\x8f\x39\x5b\xd2\x95\xe6\x89\xdd\x45\x97\xf8\xb4\xd4\x1c\x8d\xbd\xad\xe7\xd9\x69\xc2\x4e\x12\x98\x46\x10\x3c\x21\x79\x6e\x79\xab\x47\x8e\xc8\xfc\x2d\xbf\xeb\x4c\x36\x1b\x69\xd9\x40\x2f\x2e\x27\xcd\x95\x33\xb3\xe8\xfb\x8d\xa6\x5c\x77\xb7\x0e\xd4\x5e\xc7\x6d\xb4\xed\x38\x9f\xb7\x2d\xe8\xff\xed\x79\x6e\x63\x6d\x56\xda\x8b\xb9\xe0\x59\xae\xba\x0d\x5f\x51\x65\x86\xa4\xb7\xa8\x9c\x77\x72\xde\xbf\x04\x0d\x2c\x0c\xc3\x0e\x72\x43\xd2\x42\xfb\x71\x7b\xeb\xb7\x05\x24\xd3\x4d\xcc\xef\xac\x52\x96\xe0\x6d\x04\x47\xc7\xdd\x13\x5c\xc2\x32\x5e\x30\x85\x49\x04\xf5\x39\x1c\x35\xe9\xa8\xd6\x54\x86\x2b\x54\x27\x44\x91\x91\xa5\x44\xa5\x89\x6a\xcd\x93\xd6\x0c\x67\x81\x43\x9b\xd5\x1b\x9a\x11\x6b\x0e\x17\x97\xb3\x4e\x64\x42\x77\x93\xb9\x01\xb5\x45\xa6\x57\x68\x47\x46\xc1\x54\xd3\x20\x18\xf7\x3a\x5b\xa8\xd6\xc8\x46\x3b\xcb\xae\x56\xd1\xb2\x96\x55\x7e\xb4\x8e\xaa\x59\x0c\xa9\xc2\xac\xe3\x42\x43\xaf\x62\x54\x4f\xb7\x12\xec\xd1\x37\x95\xed\x29\x9b\xd5\xbe\xce\xd6\x11\x6a\x4c\x54\xbc\x6e\xc4\xda\xea\x96\xdd\x8f\xab\xb1\xba\xcc\xcc\x5c\x13\x0d\x67\x67\xd5\x99\x6a\x98\xdb\x20\x53\x5d\x73\xba\xad\x55\xd5\xd5\x0a\x2e\x77\xea\xea\x87\x79\x21\xd7\x6d\x78\xdf\xa5\xa6\x18\xe6\xe0\xfb\xb3\xfd\x13\x98\xc0\x8c\x6f\xb0\xeb\xb0\x39\x04\x83\x54\x0c\x65\x9e\xd2\x18\x4b\xd4\x04\x8e\x1f\xc8\x44\xdd\x32\x1e\x91\x0c\xab\x33\x98\x8f\x9a\x34\xad\x9c\x54\x5a\xc3\x69\xa9\xde\x26\x1e\x95\x19\x47\x00\xc3\xc9\xa9\x3d\x3b\x3c\x41\xbb\x46\xf9\x88\x0c\x55\x4a\x83\x29\x2a\x5b\x61\x35\x54\x75\x32\x55\x6b\x0f\xa7\xaa\x82\x3c\x2e\x57\xae\x58\x86\x93\xd5\xf5\xf1\xa0\x9c\x99\xb3\x99\x50\xdd\x36\xf1\x84\x92\x94\xaf\xf6\x1a\x6b\xde\xb3\x25\xc0\x5d\x8d\x1c\x59\x7f\x47\x81\xcb\x89\xb9\x28\x26\xe0\xdc\x7c\x3a\x35\xef\xed\x59\x72\xa4\xf5\x8f\xf4\x8b\xfb\xdc\x7f\x42\x93\x23\xbe\x3c\xa2\x31\x67\x47\x31\xcf\x72\xce\x90\x29\x5f\x83\xcc\x74\x7a\xa4\xf8\x10\xa6\x9f\xa5\xf2\x76\x0c\xb5\x0b\x30\x37\x9e\xcc\x06\x41\xc6\x43\x98\x97\x9e\xda\x29\xa2\x49\x9f\x72\x93\xf9\x1c\x7c\x92\xf9\x83\xdc\xb1\x9b\x99\x6b\xd3\xde\x2d\xb6\x56\x09\xb9\x93\x17\x66\xf3\xcb\x90\x74\xfc\x68\xcc\x6a\x3b\x33\xf9\x3f\x36\xd3\x0d\xb7\x37\x98\xb9\xaa\xec\xbf\x20\x49\x79\xed\xc3\x92\x8b\xea\x9e\x87\xff\x99\xa7\x71\x97\xce\xfd\x74\x3e\x15\xb8\x94\x17\x02\x97\x97\xa1\xae\xea\xc8\x4d\x18\x53\x4b\x27\x63\x9c\x74\x6f\x6c\x6a\x34\x07\x76\xe5\xec\x35\xcb\x0b\xd5\xdc\x91\xe1\xcd\x07\x1d\xcc\x60\x47\x68\x71\x64\x5f\x79\x1b\x39\xee\x71\x47\x97\xb5\xfc\xcd\xca\x18\x1c\xac\xb0\xdb\xea\x43\xd5\x7e\xd8\x6a\xbe\xd7\xea\xdf\x2c\x78\xd7\xd1\x5e\xf1\x1d\x15\xd0\xaf\x86\xcf\x59\xf2\x3e\x4f\x39\x19\x9c\xf3\xea\xc0\x76\x83\x7e\xb7\x51\xee\x9d\xce\xf2\xa2\x9e\xce\x26\x83\x3f\x77\x7c\x36\x93\x51\xb4\xbb\x7c\xb7\x0f\xfd\x1c\xd1\x42\x57\xcb\x0e\x0d\x3b\xb4\xef\x4a\xd3\x4e\xcb\xbf\x32\x2c\xee\x32\xd3\x7d\xe7\xb2\x2f\xdf\xbe\xe3\x02\x02\xd7\x8b\x4f\x17\xf0\xdf\x4e\x78\x0d\xdf\xf6\xb2\xc3\x1c\xfa\x50\x32\x12\x5f\x5f\x11\xd1\xef\x26\xbb\xdf\x06\xfe\x0a\x00\x00\xff\xff\x90\x7f\x45\x1b\x6c\x16\x00\x00")

func clientAppJsBytes() ([]byte, error) {
	return bindataRead(
		_clientAppJs,
		"client/app.js",
	)
}

func clientAppJs() (*asset, error) {
	bytes, err := clientAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/app.js", size: 5740, mode: os.FileMode(420), modTime: time.Unix(1496999441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\xdd\x8e\xdb\xb6\x12\xbe\xdf\xa7\xe0\x21\x02\x78\x17\x67\x25\x39\x69\x8b\x06\x0b\xcb\x4d\x9a\xa4\x40\x5a\x24\x0d\xda\x04\x68\xaf\x0a\x4a\x1c\xdb\xac\x29\x92\x25\x29\x79\x8d\x45\x1e\x26\xcf\x52\xf4\xbd\x0a\x52\xa2\x2d\xc9\x72\x6c\x27\xdb\x4d\x53\x5f\x99\xe4\xcc\x37\x7f\xfc\xc6\x63\x4e\xfe\xf7\xf4\xc7\x27\xaf\x7f\x7d\xf5\x0c\x2d\x6c\xc1\xa7\x67\x93\x05\x10\x3a\x3d\x43\x08\xa1\x49\x01\x96\xa0\x7c\x41\xb4\x01\x9b\xe2\xd2\xce\xa2\x87\xb8\x7d\xb4\xb0\x56\x45\xf0\x47\xc9\xaa\x14\xff\x12\xbd\x79\x1c\x3d\x91\x85\x22\x96\x65\x1c\x30\xca\xa5\xb0\x20\x6c\x8a\x9f\x3f\x4b\x81\xce\xa1\xa3\x29\x48\x01\x29\xa6\x60\x72\xcd\x94\x65\x52\xb4\xe4\x2b\xb0\xa0\x99\x20\x59\x16\x9b\x25\x02\xca\xac\xd4\x03\xca\xa4\xb4\x0b\xa9\x5b\x7a\xdf\xff\xf9\x4e\xa0\x1f\xca\xbf\xde\x09\x46\x86\x14\x2a\x06\x2b\x25\xb5\x6d\xa9\xac\x18\xb5\x8b\x94\x42\xc5\x72\x88\xfc\xe2\x12\x31\xc1\x2c\x23\x3c\x32\x39\xe1\x90\xde\x8f\xc7\x97\xa8\x60\x82\x15\x65\xb1\xdd\x0a\xe8\x96\x59\x0e\xd3\x21\x87\x27\x49\x7d\x76\x56\x0b\x72\x26\x96\x48\x03\x4f\xb1\xb1\x6b\x0e\x66\x01\x60\x31\x5a\x68\x98\xa5\x38\x49\x66\x52\x58\x13\xcf\xa5\x9c\x73\x20\x8a\x99\x38\x97\x45\x92\x1b\xf3\xcd\x8c\x14\x8c\xaf\xd3\x9f\x64\x26\xad\xbc\xfa\x62\x3c\xbe\xfc\x72\x3c\xbe\xfc\x6a\x3c\xbe\xfc\xba\xfe\xce\x2c\xe1\x2c\x0f\xee\x7c\x80\x15\x96\x4b\x11\xcc\xbc\x20\x2e\x0e\xc2\xff\xff\x3c\x97\xc2\x1c\x03\xea\x2e\x80\xb9\x4a\x92\x52\xa8\xe5\xdc\xe3\x55\x25\x44\x45\x83\x93\x50\x66\x6c\x67\x27\xce\x8d\xc3\xad\x81\xeb\xda\x23\xa3\xf3\x3d\x40\x1b\xfd\xf8\x77\x83\xa7\x93\xa4\x56\x98\x1e\xa9\xbd\x31\xfa\x01\xaa\xec\xfe\x43\xb1\xf5\xde\xad\x4e\x77\x81\x5c\x33\x69\x6a\x10\xff\x35\x2e\x58\x1f\x65\x92\xd4\x64\x9b\x64\x92\xae\xa7\x67\x13\xca\x2a\xc4\x68\x8a\x89\x52\x9b\xfb\x4b\x23\x2b\x25\xcf\x88\x46\x39\x27\xc6\xa4\xb8\xa0\x11\x05\x61\x02\x9f\xbc\xd4\xe2\x41\xeb\xd4\x5f\x3c\x8c\x7c\xb5\x52\x3c\xe3\x70\x7d\x85\xee\xe3\xee\x2d\x9d\x24\x8b\x07\x2d\xfd\x82\x46\x59\x69\xad\x14\x2d\x18\x77\x33\x9a\x5d\x8c\x1e\xe5\x9c\xe5\xcb\x58\x10\xcb\x2a\x48\xb1\x21\x15\x3c\x16\xf4\x8d\xe2\x92\xd0\x96\x23\x01\xcc\xe9\x4e\x33\x92\x2f\x4b\x35\x49\xc2\x7a\x6b\x2e\xd9\xd8\x6b\xa2\x4c\xb6\x61\x36\x3b\x2e\x15\xc1\x17\xc2\x44\xd4\x50\x36\x5c\x9e\x60\x88\x32\xc2\xe5\x3c\x52\x5a\x16\xca\x76\xdc\x70\x9f\xab\x90\x8d\x14\xd7\x12\x71\x9d\x9b\x41\x41\x26\x54\x69\x23\xc5\x49\x0e\x0b\xc9\x29\xe8\x8d\x52\x6b\x6f\x40\xb5\x22\xbc\xdc\x1a\xf0\xab\x5d\xa9\x47\x1e\x3d\xc5\x52\x3c\x77\x5f\x76\x05\x3c\x9d\x6a\x8c\xa7\x3e\x26\xdc\xcb\x57\x27\xd2\x5e\x16\x2c\xc9\x0c\x2a\x68\x34\x63\xd7\x40\x77\xab\x61\x49\xe6\xaf\x95\x80\x95\xc1\xa8\x8a\x32\x26\xa8\x8b\x98\x93\xcc\xd1\xfa\x9e\x3d\x1f\x15\x60\x0c\x99\x43\xec\x44\x46\x17\xbd\x8a\x6e\xaa\xea\x33\xe4\x2a\x41\x98\x00\xbd\x2b\xe4\x05\x3d\xea\xf4\xe6\x06\xdd\xb3\xe7\xb8\x8d\xfb\x1a\xae\x2d\xbe\x40\x6f\xdf\x4e\x92\x5a\x66\x58\x3f\x18\x0a\x8e\x76\x2a\xd2\xf7\xf5\xd5\xf6\x70\x74\x81\xd1\xa3\x25\xac\x4b\x15\x83\xb0\xa0\x37\x97\x55\x8a\x97\xb0\x7a\xe9\x62\x1f\x34\xb8\xfd\x54\x51\x21\xa9\xcb\x88\x68\xe4\xa7\xf5\xe5\x75\xde\x0c\x64\x24\x19\x4a\xc9\x60\xe2\x38\x33\x03\x00\xed\xd3\x88\x59\x28\x50\x15\xcd\xa4\x4e\xf1\xb9\x5b\xb8\x1f\x22\x0a\xd7\x17\x88\x09\xd4\x29\xdc\x12\xd6\x29\x76\x12\x03\x55\xea\x24\xd1\x71\x8e\x89\x99\xfc\x4d\x96\x96\x33\x01\x03\x4c\xdc\xd1\x33\x8a\x88\x29\xba\xb9\x41\xde\x1f\x57\x2a\xbf\xf3\x5e\x43\xef\xeb\x1c\x28\xc4\x47\xf2\xfa\x57\xbe\xd7\x48\x34\x14\xb2\x02\x97\xec\xf3\x3a\xdc\xf7\x04\xd5\x0e\xac\x65\x4d\x69\x56\x10\xbd\xc6\x53\x0a\x1c\xec\x51\x51\xf6\xfb\xcf\xa0\xc0\xa6\x2c\x7b\x4a\xbf\x5b\xd4\xba\x8d\x91\xac\x77\x0b\xda\x14\x34\xa0\xdd\xac\x71\x88\x86\x41\xec\x9f\xa0\x62\x83\x7d\xdb\x6c\x6c\x60\x8f\x26\xe4\xcf\xb5\xfc\x29\x9c\x0c\x2a\xff\x16\x5a\xee\x14\xf3\x4e\xa9\x89\xee\x86\x9b\x4d\xd2\xff\x53\xf4\xec\xbb\x1c\xd8\xb9\x90\xa5\x3e\x44\x4d\x2f\xb3\x97\x97\x96\x64\x6e\xe8\xdf\x97\x1b\x7f\x1c\xb9\xa1\x6f\x1f\x67\xbb\x92\x5a\xae\x0e\x27\x7c\x0b\xda\xa7\x3a\x25\xeb\x86\xe2\x5d\xb1\x8f\x80\x24\xc5\x6d\x23\xaa\x13\x11\x5b\x82\x7b\xd3\xd3\x03\x1b\xec\x00\x5d\xcf\xea\x09\xfc\x70\x04\x5a\xae\x36\xfd\x40\xcb\x55\xbb\x1d\xf8\x9b\xe1\x72\x6e\x30\x6a\x5a\x81\x3b\x3b\x82\x32\x35\x74\x0e\xdc\x37\x6b\x2d\x57\x0e\xa5\x9b\x13\x7f\x78\x0a\xd0\x81\xbe\x1a\x1a\x4a\x63\x8f\x1c\xf3\x63\xdf\xb6\x74\xd2\xdf\x05\xa9\x40\xd4\x53\xed\xf9\xa8\x3d\xe3\x8e\x2e\xd1\x88\x14\xa3\x90\xc4\x03\x99\x6a\x9b\xf7\x1d\xc4\xfd\xd3\x3e\xa2\x9f\x6c\xf4\x0e\xf5\x95\x8e\xe0\x9d\xa4\x5d\x7d\xb2\xb4\xab\xcf\x38\xed\x27\x76\x80\x61\x6a\x6f\x05\xf6\x4c\x07\x1f\x31\x62\xcd\xa4\xb4\x42\xda\x5b\x9f\xb1\x02\xee\xd1\x43\xd6\x77\x8d\xc2\x29\x53\xd6\x46\xe7\x36\xc7\x2c\x14\x5e\x2d\x3e\x66\xde\xaa\x1b\x6c\x48\xc2\x89\x63\xd7\x27\x1a\x9f\x42\x36\x3f\xd3\xf9\xe9\xc0\x00\x15\x76\x4c\x78\xc7\xa1\xac\x0a\xcf\x7b\x05\x8d\x8c\x20\xf9\x32\x23\xba\x7e\xdd\x08\xab\xf6\xcb\xc6\xa6\x19\x82\x51\x2f\xcc\x7c\xa0\x1b\xb6\xca\xd0\x4b\xf0\x3d\x0d\x33\x13\x07\xd4\x38\xe7\xd2\xc0\xf9\x05\xee\x93\xd1\xef\x7f\x6b\xc5\x76\xca\xd8\x79\x7a\x0a\x18\xd3\xb3\x26\x82\xf0\xb4\x47\x61\x06\xba\x7e\xe0\x23\x4a\xed\xbc\xde\x35\xcf\x76\x89\x7f\x39\xff\x3b\x00\x00\xff\xff\xfc\x47\x14\xcd\x49\x17\x00\x00")

func clientIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_clientIndexHtml,
		"client/index.html",
	)
}

func clientIndexHtml() (*asset, error) {
	bytes, err := clientIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/index.html", size: 5961, mode: os.FileMode(420), modTime: time.Unix(1496998880, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"client/app.js": clientAppJs,
	"client/index.html": clientIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"client": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{clientAppJs, map[string]*bintree{}},
		"index.html": &bintree{clientIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

