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

var _clientAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\xcd\x6e\xdb\xb8\x13\xbf\xfb\x29\x06\x6a\x01\xd9\xf8\xdb\x32\x72\x2b\x64\xf8\xd0\x36\xfd\x63\x8b\x45\xe2\x22\xfd\x00\xb6\x41\x0e\x8c\x34\xb6\x89\x48\xa4\x96\xa4\x9c\xa4\x85\xef\x7b\xe9\x43\xf4\xb0\x87\x3d\xec\x2b\xf4\xd2\xe6\xbd\x16\xa4\x28\x59\x1f\x94\xed\xec\x2e\xb0\x3e\x24\x32\xe7\x37\x9c\xaf\x1f\x87\x23\x7b\xb9\x44\x90\x4a\xd0\x48\x79\xb3\xc1\xe0\x43\x8e\x41\x2e\x71\xf8\x21\xc7\x33\xa2\x50\x50\x92\x8c\x66\xf5\xd5\xd7\x27\xcf\xd8\x68\x36\x18\x44\x9c\x49\x05\xf4\xe4\x19\x83\x39\x30\xbc\x05\x2b\x1b\x7e\x1e\x00\x00\x24\x3c\x22\x09\x86\xe0\xcb\x1b\x7f\x0c\x30\x9d\xc2\xbb\xc5\xe9\x22\x04\x89\xca\xca\x60\x29\x78\x0a\x51\x42\x91\x29\xa3\x92\xa2\x94\x64\x85\x32\x84\x62\x0b\xfd\x41\x56\xff\x56\x43\xb5\x97\xf5\x87\xe1\xad\x0c\xc1\x3f\xc7\x5b\xe9\x8f\x9d\xd2\x77\x78\xa7\x2c\x02\x14\xde\xa9\x1e\xd8\x9b\x84\x44\xb8\xe6\x49\x8c\x22\x04\xff\xdd\x7d\x86\x66\x19\x08\x8b\x21\x13\x28\x25\x20\x53\x28\x1c\xda\x12\xc5\x86\x46\x3a\x06\xff\xad\x7d\xec\x47\x59\x6f\x2c\xb0\xcf\x21\x0b\x76\xf8\x64\x25\x47\xb8\xb5\xe6\xb9\xd0\x3e\x2d\x32\x64\x94\xad\x8a\xef\x0e\x5c\x4c\xee\x43\xf0\x4f\xc9\xbd\x43\x46\xd2\x10\xfc\xe7\x67\x0e\x49\xa6\x25\x6f\x5c\x92\x25\xe7\x8a\x71\x85\xbb\xbc\x57\x4b\x7d\xe1\x96\x72\x47\xbc\x95\xea\xe1\x80\xa3\x84\x4b\x7c\xa1\x58\x08\xfe\x4b\xfd\xe8\xca\x2b\xd9\x60\xbc\xb8\xd1\x05\xd0\x4f\xb0\xf8\xd9\x01\x42\x21\xce\xf9\x05\xca\x8c\x33\xa9\xd9\xfc\xea\xe2\x62\x71\x11\xc2\x39\x07\x61\x17\xfd\x86\xce\xb6\xfa\xb6\xdd\x6d\x26\x6f\x1e\x4f\xe2\xc5\x27\x46\x52\x57\x19\x6a\x34\xd6\xff\x80\x6b\x60\x7e\x14\x91\xcf\x49\xf6\xfd\xcf\x87\xaf\x85\x0a\x10\x90\x2a\x21\x3f\xbe\x1c\xc5\xe5\x24\x7f\xf8\x76\xed\x72\xa7\x49\xe5\x8f\x0c\x19\x45\x90\x87\xe0\x6e\xbf\x0a\xb5\xfc\xb0\x67\x15\x9d\x45\x4c\x19\xf9\xf1\x85\x7d\xff\x03\xd6\x3c\xa6\xcc\x65\xd2\x72\x1a\x1f\x7e\xeb\x23\xf5\x29\xe7\xd7\x18\xbb\x72\x58\x30\xdb\x88\x5d\x14\x6a\xd1\xdb\xd4\x23\xe3\x9f\xd8\xf7\xaf\xe9\x8d\xcb\x15\x37\xb7\xcb\xf0\x4b\xcd\x23\x12\x50\xa3\xf7\x47\xa2\x36\x5c\xd0\x87\xdf\xf7\x52\xfc\x7d\xc2\x1f\xbe\xa1\xce\xd3\x51\x34\x7f\xf9\xd3\x2f\x2f\x9e\x87\xa6\x5a\x28\x80\x21\x8f\x33\xbe\xc1\x98\x24\xbd\x6c\x1f\x14\x7f\xb7\xfa\x6a\x18\x4c\xa7\xb0\x49\x81\x4a\xd8\xe4\x08\x94\x49\x45\x98\x6d\x53\x28\x04\x17\x5a\x52\x3c\x08\x94\x79\xa2\x40\xa0\xca\x05\xc3\x18\xae\xef\x81\xdc\x51\x2e\x07\xcb\x9c\x45\x8a\x72\x06\x72\xcd\x6f\x5f\x69\xec\x70\x93\x8e\x0b\xad\x91\x3d\x35\x74\x09\x43\xb3\x10\x94\x67\x71\x54\x3b\x4f\xfa\x7e\xe2\x09\x06\x09\x5f\xb5\x50\x41\x4c\x14\x19\xcd\x8e\x41\x4a\x45\x54\x2e\x8f\xc3\xae\x91\xc4\x28\xea\xe0\x4d\x6a\xa4\x67\x72\x05\x73\x70\xf8\x50\x20\xb7\x80\x89\xc4\x46\x34\xbf\xe6\x28\xd5\xa1\x60\x0a\x50\x8f\xb5\x4d\x1a\x3c\x55\x43\xcf\xf6\x99\xa0\x51\x5f\x6f\xd4\x30\xec\xb6\xe2\x9b\xa4\xfb\x36\xe5\x81\xdd\x68\x7f\x70\x16\x64\x77\x1f\xb8\x1d\x8f\x38\x5b\xd2\x95\xa6\x89\xdd\x45\x57\xf8\xac\xd0\x1c\x8e\x06\xdb\xc1\xc0\x0e\x13\x76\x90\xc0\x24\x04\xff\x09\xc9\x32\x4b\x5b\x3d\x71\x84\xe6\x6f\xf1\x5d\x27\xb2\xde\x47\x8b\xfe\x79\x79\x35\xae\xaf\x9c\x9b\x45\xcf\xab\xf5\xe4\xaa\xb9\xb5\xa0\xf6\x36\x6e\xa2\x6d\xc3\xf9\xbc\x6d\x40\xff\x6f\x8f\x73\x13\x6b\xb3\xd2\x5c\xcc\x04\x4f\x33\xd5\xee\xf7\x8a\x2a\x33\x23\xbd\x45\xe5\xbc\x92\xb3\xee\x1d\x68\x60\x41\x10\xb4\x90\x1b\x92\xe4\xda\x8f\xbb\x3b\xaf\x29\x20\xa9\xee\x61\x5e\x6b\x95\xb2\x18\xef\x42\x98\x9c\xb4\x0f\x70\x01\x4b\x79\xce\x14\xc6\x21\x54\xc7\x70\x58\x67\xa3\x5a\x53\x19\xac\x50\x9d\x12\x45\x86\x25\x99\xac\x26\xaa\x35\x8f\x1b\x23\x9c\x05\xf6\x6d\x56\x6d\x68\x26\xac\x39\x5c\x5e\xcd\x5a\x91\x09\xdd\x4c\xe6\x06\xd4\x14\x99\x56\xa1\x1d\x19\xfa\x53\x4d\x03\x7f\xd4\x69\x6c\x81\x5a\x23\x1b\xee\x2c\xbb\x3a\x45\xc3\x5a\x5a\xfa\xd1\x38\xa9\x66\x31\xa0\x0a\xd3\x96\x0b\x35\xbd\x92\x51\x1d\xdd\x52\x70\x40\xdf\x54\xb6\xa3\x6c\x56\xbb\x3a\x5b\x47\xa8\x11\x51\xd1\xba\x16\x6b\xa3\x59\xb6\x3f\xae\xbe\xea\x32\x33\x73\x0d\x34\x9c\x9d\x97\x67\xaa\x66\x6e\x83\x4c\xb5\xcd\xe9\xae\x56\x56\x57\x2b\xb8\xdc\xa9\xaa\x1f\x64\xb9\x5c\x37\xe1\x5d\x97\xea\x62\x98\x83\xe7\xcd\x0e\x0f\x60\x02\x53\xbe\xc1\xb6\xc3\xe6\x10\xf4\x52\x31\x90\x59\x42\x23\x2c\x50\x63\x38\xd9\x93\x89\xaa\x65\x3c\x22\x19\x56\xa7\x37\x1f\x15\x69\x1a\x39\x29\xb5\xfa\xd3\x52\xbe\x4c\x3c\x2a\x33\x8e\x00\xfa\x93\x53\x79\x76\x7c\x82\x76\x8d\xf2\x11\x19\x2a\x95\x7a\x53\x54\xb4\xc2\x72\xa6\x6a\x65\xaa\xd2\xee\x4f\x55\x09\x79\x5c\xae\x5c\xb1\xf4\x27\xab\xed\xe3\x51\x39\x33\x67\x33\xa6\xba\x6d\xe2\x29\x25\x09\x5f\x1d\x34\x56\xbf\x67\x0b\x80\xbb\x1a\x19\xb2\xee\x8e\x02\x97\x63\x73\x51\x8c\xc1\xb9\xf9\x74\x6a\x5e\xdb\xd3\x78\xa2\xf5\x27\xfa\xbd\x7d\xee\x3d\xa1\xf1\x84\x2f\x27\x34\xe2\x6c\x12\xf1\x34\xe3\x0c\x99\xf2\x34\xc8\x0c\xa7\x13\xc5\xfb\x30\xdd\x2c\x15\xb7\x63\xa0\x5d\x80\xb9\xf1\x64\xd6\x0b\x32\x1e\xc2\xbc\xf0\xd4\x4e\x11\x75\xfa\x14\x9b\xcc\xe7\xe0\x91\xd4\xeb\xe5\x8e\xdd\xcc\x5c\x9b\xf6\x6e\xb1\xb5\x8a\xc9\xbd\xbc\x34\x9b\x5f\x05\xa4\xe5\x47\x6d\x54\xdb\x99\xc9\xfe\xb1\x99\x76\xb8\x9d\xc1\xcc\x55\x65\xef\x05\x89\x8b\x6b\x1f\x96\x5c\x94\xf7\x3c\xfc\xcf\x3c\x8d\xda\x74\xee\xa6\xf3\xa9\xc0\xa5\xbc\x14\xb8\xbc\x0a\x74\x55\x87\x6e\xc2\x98\x5a\x3a\x19\xe3\xa4\x7b\x6d\x53\xa3\xd9\xb3\x2b\x67\xaf\x59\x96\xab\xfa\x8e\x0c\x6f\x3f\xe8\x60\x7a\x3b\x42\x83\x23\x87\xca\x5b\xcb\x71\x87\x3b\xba\xac\xc5\x4f\x56\xc6\x60\x6f\x85\xdd\x56\xf7\x55\x7b\xbf\xd5\xec\xa0\xd5\xbf\x59\xf0\xb6\xa3\x9d\xe2\x3b\x2a\xa0\xdf\x0c\x9f\xb3\xf8\x7d\x96\x70\xd2\x3b\xe7\x55\x81\xed\x06\xfd\x76\xa3\x3c\x38\x9d\x65\x79\x35\x9d\x8d\x7b\x7f\xed\xf8\x6c\x26\xa3\x70\x77\xf9\x6e\xf7\xfd\x1a\xd1\x40\x97\xcb\x0e\x0d\x3b\xb4\xef\x4a\xd3\x4c\xcb\xbf\x32\x2c\xee\x32\xd3\x7e\xe5\xb2\xef\xde\x9e\xe3\x02\x02\xd7\x8b\x4f\x1b\xf0\xdf\x4e\x78\x35\xdf\x0e\xb2\xc3\x1c\xfa\x40\x32\x12\xdd\x5c\x13\xd1\xed\x26\xbb\x9f\x06\xfe\x0a\x00\x00\xff\xff\x43\x4f\x79\x3f\x6c\x16\x00\x00")

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

	info := bindataFileInfo{name: "client/app.js", size: 5740, mode: os.FileMode(420), modTime: time.Unix(1497000378, 0)}
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

