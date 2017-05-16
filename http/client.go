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

var _clientAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x8a\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\xb7\x2e\x46\xc1\x13\xa4\xf7\xe8\x3e\xb5\x9f\x30\x20\x63\x99\x4c\x70\x21\xbd\xbb\x98\xbe\xdd\xe3\xbd\xa9\x37\xa2\x85\xeb\x33\xa6\x39\x25\xe3\x07\x4b\xe7\xed\x9b\x00\x80\xaf\x0c\xb9\x96\x7d\x97\xc7\xf0\xad\x44\xc9\x38\xdb\x9f\xea\x64\xa8\xd5\x0c\x59\x18\xeb\x0a\x6e\x1a\x6f\x87\x36\x78\x37\x53\xab\x17\x19\xf3\x91\x8e\xfb\xfc\x0b\x00\x00\xff\xff\xca\xcf\x8b\x2c\x6b\x00\x00\x00")

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

	info := bindataFileInfo{name: "client/app.js", size: 107, mode: os.FileMode(420), modTime: time.Unix(1494930498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xbf\x4e\xc3\x40\x0c\x87\xf7\x7b\x0a\x73\x0f\x50\xab\x2b\x72\xbd\x00\x33\x0c\x08\x89\x31\x89\x4d\xce\xd0\x3f\xa7\x3b\x37\x12\xaa\xf2\xee\x28\x5c\x32\x75\xb2\x65\xeb\xf7\x7d\x36\x3d\x3c\xbf\x3e\xbd\x7f\xbe\xbd\x40\xf2\xd3\x91\x03\x6d\x45\x3b\xe1\x00\x00\x40\x6e\x7e\x54\xfe\x50\xef\x7b\x15\x73\xc2\x36\x68\xcb\x3a\x14\xcb\x0e\xb5\x0c\x87\x98\xdc\x73\x7d\x44\xbc\x9e\xf3\xcf\xb8\x1b\x2e\x27\x9c\xae\x8a\x62\xd5\x97\x66\xf7\x5d\x23\x13\xb6\x00\x07\xc2\xa6\xa0\xfe\x22\xbf\x1c\x48\x6c\x02\x93\x43\xec\x72\x8e\x2b\x3b\xed\xf9\x76\x83\xb1\xa8\xba\x9d\x47\x98\x67\xc2\xb4\x5f\x92\x62\x13\x87\x4d\x2d\xfa\xa5\xa5\x1d\xd0\xe5\x7c\x67\x59\xf1\xf8\xff\xd7\x5f\x00\x00\x00\xff\xff\x83\xca\xcd\x26\xee\x00\x00\x00")

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

	info := bindataFileInfo{name: "client/index.html", size: 238, mode: os.FileMode(420), modTime: time.Unix(1494930732, 0)}
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

