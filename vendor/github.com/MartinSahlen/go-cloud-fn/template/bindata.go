// Code generated by go-bindata.
// sources:
// index.js
// DO NOT EDIT!

package template

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

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4f\x6f\xdb\xc8\x0f\xbd\xfb\x53\xb0\xbd\x48\x42\x1d\xa9\x97\x1f\xf0\x43\x02\xef\x02\x05\x5a\x74\x17\xdb\x6e\xd1\x64\xb1\x87\xa2\x28\x26\x1a\xda\x9a\x66\x34\x54\x39\x54\x1c\xc3\xf0\x77\x5f\xcc\x1f\xcb\x72\x9a\x1c\x16\xd8\x53\xe2\x21\xf5\xf8\x48\x3e\x72\xa6\x69\xde\x2b\xa7\x2d\xc2\x1b\xd5\xde\x6d\x98\x46\xa7\x01\xef\xd1\x89\x07\xd5\xb6\xc4\xda\xb8\x0d\x08\x81\x1f\xb0\x5d\xac\x47\xd7\x8a\x21\x07\xbe\x33\x7d\xfa\x8e\x4b\xad\x44\x55\xb0\x5f\x00\x30\xca\xc8\x0e\x1c\x6e\xe1\x13\x53\x6f\x3c\x96\x25\xa3\x27\x7b\x8f\x4b\x60\xfc\x8e\xad\x54\xb0\xfa\x25\xfa\x02\x34\xcd\xf5\xa0\xb6\x0e\xa4\x43\x98\x80\x95\xd3\x60\x5c\xf0\x8c\xe7\xe8\xee\x61\xcd\xd4\xc7\x1f\x83\x62\x74\x02\x03\x53\x8b\xde\xd7\x11\xa4\x25\xe7\x05\x06\x58\x01\xe3\x8f\xd1\x30\x96\x45\xdb\x19\xab\xbf\x65\xaf\xa2\xaa\xf1\x01\xdb\x77\xc6\x62\x59\xd4\xcd\x7e\x5f\xbf\xcb\xa1\x3e\xaa\x1e\x0f\x87\x62\x09\x5f\xbe\x2e\x33\x23\x08\xf1\x2e\xa7\x00\xe8\xee\x97\xf1\xfc\x50\x5d\xc5\xbf\xf7\x8a\xc1\x2a\x2f\x1f\xd0\x7b\xb5\xc1\x74\x38\xd4\x5e\xb4\x71\xb5\x47\x79\xeb\x5a\x0a\xf5\x2a\x8b\x51\xd6\x17\xff\x2f\xf2\x67\x4d\xf3\x07\x6d\xc0\x8b\x72\x5a\xb1\x06\x64\x86\x3e\x21\xf8\x58\xd9\x99\xe1\x84\x88\xcc\x35\xb9\xb2\x08\xd5\x2d\x96\x50\x22\xf3\xac\x76\x29\x71\xb2\x58\x23\x33\x71\xb0\xd6\x42\xd7\xc2\x21\x7a\x95\xe3\x1e\xaa\x13\x1c\x8d\x72\x06\x47\xa3\x3c\x09\x67\x69\x13\x6c\x3f\x83\xc1\x3c\x73\x58\x01\x8d\xf2\x28\x4a\x80\x6f\x2d\x79\x0c\xf8\x2d\x69\x3c\x0b\x60\xd6\xe9\x10\x5e\xac\x56\xf0\xba\x9a\xce\x43\x79\x6e\x3a\xe3\xa1\x47\xe5\x7c\xec\x73\x10\x17\xac\x95\xb1\xa8\xa1\x81\x41\x39\xd3\xde\xa1\xae\xe1\x9a\x60\x8b\x59\x47\xd0\x29\xd6\xf5\x84\x91\x0e\xcb\x89\xeb\x01\xd0\x7a\x3c\x0b\x02\x9f\x93\x12\x93\x94\x92\x3c\x61\x6b\xa4\x8b\x07\x56\x09\x7a\x09\x59\x0d\xa3\x24\xc9\xa5\xaa\xcd\x11\x7e\x73\xd0\x2a\x8f\x40\xeb\xc8\xb1\x0f\xa3\xd1\x89\x0c\x4b\x90\x90\x81\x49\xf4\x19\xfd\x40\x2e\xb8\xdd\x06\x52\x73\x92\x91\x40\x39\x2b\xe4\x89\xf0\x99\xd0\x9a\xe6\x6f\x36\x92\xa8\x26\x94\x26\x2b\xa6\x09\x32\x0f\x4c\x85\xa6\x5a\x15\x1e\xa2\x04\xe3\xf0\x78\xb3\x71\xca\x66\x94\xb7\x4e\x07\xb2\xc6\x0d\x63\xe6\x71\x54\xeb\x36\xe0\x97\xbf\x5f\xff\xf9\xb1\xf6\xb1\xd1\x66\xbd\x4b\x93\x5c\x9d\xcb\x1a\x9d\x4e\x65\x0d\xdc\x0e\x8b\xc5\xb4\x2f\x42\xe2\x90\xd9\x9c\x16\x43\x17\x8d\xef\x45\x86\x92\xf1\x47\x98\x7a\x9f\x7a\x1d\x66\x27\x7b\xbf\x21\xbd\xbb\x5a\x2c\x00\xfc\xd6\x48\xdb\x41\xf0\xac\x37\x28\x65\xd1\x92\x13\x74\x72\x21\xbb\x01\x8b\xea\x28\x92\x58\xf3\x42\x0d\x83\x35\xad\x0a\x51\x9a\xef\x9e\x5c\x71\xb9\x38\x96\x75\x42\x85\x15\x3c\x4a\x29\x40\xdf\x92\xde\x4d\x85\xbe\x65\x54\x77\x57\xcf\xe0\x3e\x5c\x6c\xb7\xdb\x8b\x35\x71\x7f\x31\xb2\xc5\x30\xce\xa8\xa7\x40\x41\xa8\x08\x01\x2d\xac\x22\x8f\x0c\x6b\x62\x68\x2d\x8d\x7a\x5a\x60\x1e\x34\x85\xb1\xee\x8c\x5f\x82\x27\xf8\x3e\x86\xf5\x64\xd5\x0e\x94\x25\xb7\x99\x90\xa2\xf0\x8c\x04\x1f\xe6\x1d\xf4\xca\xbd\x80\x0f\x6a\x77\x8b\x41\xe2\xbe\xa3\xd1\xea\xb4\xdb\x78\x6c\x05\x3c\xf5\x08\x77\x26\x75\x73\x60\x1a\x90\x27\xa4\xc0\xf6\x58\x83\x48\xee\x57\x20\x06\x47\x52\x83\x45\x29\x3c\xdc\x21\x0e\x60\x24\x69\x74\xab\x76\x91\xb5\xa3\xed\x12\x94\x5f\x9c\x8d\xa0\xf1\xd0\xd1\xf6\xa7\x8c\x6e\xb1\x53\xf7\xe8\xeb\x53\xbd\xc3\xb2\x7b\xd4\xad\xe5\x13\x1d\xaa\xfe\xfb\x0e\x51\x2b\x28\x17\x5e\x18\x55\xff\x8c\x02\x8e\x80\xcf\xe2\x09\x3e\x48\x33\x58\x65\x9e\xd3\xd0\xb3\x08\x87\x45\x56\xf2\x7a\xb4\xf6\x2f\xb6\xd9\x77\x60\x12\x6a\xc9\xc2\x2b\x28\x2e\x9b\xa6\x80\x57\x30\x49\xba\x23\x2f\x45\x95\x4f\x88\xcd\xc6\x38\x15\x3e\xbd\x3a\x42\x85\x39\xfa\x9c\xdb\xb7\xca\x8a\x2f\x42\xf4\xe2\x72\x4e\x2b\xdd\x44\x45\x87\x4a\x23\xfb\x64\xab\xf3\xaf\x6c\xeb\x51\x3a\xd2\xd9\x94\x7e\x64\x0b\x63\x4f\x82\xdf\x94\xd6\x9c\xcd\x66\xc8\xa6\x91\x6d\x71\x79\xcc\x27\xa4\x98\x26\x73\x76\xbf\xcf\x08\x86\x86\xd6\xd2\xa1\x8b\xf7\xfa\x68\xe7\x77\x48\x58\x1e\xc7\xee\xc6\xf1\x38\xba\xa4\x2a\x32\xfa\xda\x8b\x92\xd1\xc7\x35\x93\xff\xff\x16\xef\x89\x99\x07\x4a\x32\xe7\xcc\xaa\x99\xc5\xe9\x64\x9a\xc4\x12\x6f\x9e\xba\x55\xd2\x76\x65\x39\x63\x32\x0b\xf5\xbf\xd7\xaf\xab\xf9\x0a\x4b\x1b\x6c\xbf\x0f\xd7\x51\x7d\xc3\x66\xb3\x41\x7e\x7f\x73\xf3\x09\x0e\x87\x05\x3e\x0c\xc4\xe2\xbf\x14\x3f\x3f\x12\xbe\xc2\x6a\x9a\x87\x47\x7b\x2d\xbf\x79\x9e\xda\x7b\x57\x8b\x43\xd3\xc0\x7e\x9f\xee\xa2\x7f\x11\x21\xbe\xbf\x96\xd0\x2a\x6b\x6f\x55\x7b\x77\x16\x68\xde\x99\xe8\x57\xc7\xad\x9d\xba\x32\x21\x9c\x56\x67\x82\x38\xe6\x9f\xab\xf5\xbc\x5f\x78\xbb\xbd\x8d\x4f\x8a\x97\x47\x86\xf9\x2a\x7e\x59\x9d\xae\x81\x9c\x96\xd3\x21\xab\x7f\x02\x00\x00\xff\xff\xe8\x09\x2d\xf8\x42\x0a\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2626, mode: os.FileMode(420), modTime: time.Unix(1492439855, 0)}
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
	"index.js": indexJs,
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
	"index.js": &bintree{indexJs, map[string]*bintree{}},
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
