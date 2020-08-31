// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// conf/config.yaml
// locales/error.go
// locales/errors_en.json
// locales/errors_tw.json
// locales/errors_zh.json
package assets

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _confConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\xc1\x4a\xc4\x40\x10\x44\xef\xf3\x15\x7d\xda\x9c\xb2\x64\x09\xac\xda\x90\x3f\xf0\xb8\x77\xe9\xcc\x74\xd6\xc1\x99\xe9\xb1\x3b\x41\xfd\x7b\x99\x20\xe2\xea\xd1\x5b\xf1\xa8\x2a\xaa\xf2\x87\xbd\x26\x74\x00\x9b\xb1\x16\xca\x8c\xd0\xa9\xc8\xda\x39\x80\x4a\x66\x6f\xa2\xe1\x07\xa2\x10\x94\xcd\x10\xba\x3d\x78\x9c\xc9\xa2\x7f\xa2\x5a\x8f\x41\xfc\x0b\x2b\x8e\xe3\x70\x6e\xc6\x40\x2b\xcd\x64\xad\xee\xdb\xd3\xb8\x97\xb2\xc4\x2b\x42\xe7\x9f\x49\x8d\xd7\x69\x5b\x97\xfb\x43\x6d\xfa\x12\x33\x4f\x17\xdd\xf8\x90\xc4\x4f\x8f\xe2\x29\xb5\x44\xa6\xf7\x3e\x86\xc4\xbd\x97\x52\x0c\xe1\x34\x7c\x41\xa9\x5c\x6e\x60\x92\x6b\x9f\x25\x30\xc2\x42\xc9\xd8\x39\xe5\x10\x0d\x6f\x66\xef\xe8\xef\xec\xf3\x78\xf7\xf0\xeb\xf2\xfe\x62\x46\x18\x9c\xb1\x59\x94\xf2\xdf\xa2\xd3\x67\x00\x00\x00\xff\xff\x0c\xa7\xf1\x91\x6c\x01\x00\x00")

func confConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_confConfigYaml,
		"conf/config.yaml",
	)
}

func confConfigYaml() (*asset, error) {
	bytes, err := confConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/config.yaml", size: 364, mode: os.FileMode(420), modTime: time.Unix(1598752875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesErrorGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\xc9\x4f\x4e\xcc\x49\x2d\xe6\xe2\x2a\xa9\x2c\x48\x55\x70\x2d\x2a\xca\x2f\x52\x28\x2e\x29\xca\xcc\x4b\xe7\xe2\x4a\x2b\xcd\x4b\x56\xd0\x48\x2d\x2a\x82\x88\x6b\x42\x28\x0d\x4d\xa8\x02\x85\x6a\x2e\xce\xa2\xd4\x92\xd2\xa2\x3c\xa8\x00\x48\xa9\x26\x57\x2d\x17\x20\x00\x00\xff\xff\x30\x61\xfa\x20\x5c\x00\x00\x00")

func localesErrorGoBytes() ([]byte, error) {
	return bindataRead(
		_localesErrorGo,
		"locales/error.go",
	)
}

func localesErrorGo() (*asset, error) {
	bytes, err := localesErrorGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/error.go", size: 92, mode: os.FileMode(420), modTime: time.Unix(1598539326, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesErrors_enJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\x3d\x0e\xc2\x30\x0c\xc5\xf1\xbd\xa7\x78\xca\xd2\x85\xa1\x85\xad\x97\xb1\xac\xca\x88\x48\xf9\xc2\x0e\x43\x85\xb8\x3b\x4a\x02\x52\xd7\xbf\xdf\xcf\xef\x09\x70\xeb\xb2\xac\x94\x32\x3d\x5f\xa2\x07\x15\x56\x8e\x6e\x83\x2b\x41\xd8\x04\x85\xcd\x50\x1f\x82\xc4\x51\x36\xcc\x7d\x35\x63\xcc\x2e\x3f\x7f\x1d\x8c\x7c\xda\x73\x2c\x41\xaa\xb4\x0f\x4d\xf5\x0e\x6f\x38\x9d\xfe\xea\x46\xbb\x0a\x57\xa1\x3b\xfb\xd0\x80\x65\xd5\x03\x23\xa2\xc7\xe9\xf3\x0d\x00\x00\xff\xff\xad\xe8\x5c\x7f\xa3\x00\x00\x00")

func localesErrors_enJsonBytes() ([]byte, error) {
	return bindataRead(
		_localesErrors_enJson,
		"locales/errors_en.json",
	)
}

func localesErrors_enJson() (*asset, error) {
	bytes, err := localesErrors_enJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/errors_en.json", size: 163, mode: os.FileMode(420), modTime: time.Unix(1597372534, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesErrors_twJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x32\x34\x30\x30\x8c\xcf\xcb\x8f\x2f\x2c\x4d\x2d\xaa\x8c\x2f\x48\x2c\x4a\xcc\x55\xb2\x52\x50\x7a\xb1\xba\xfb\x69\xd3\xe6\x97\x8d\x53\x9e\xcd\x5f\xfa\x62\xfd\xa2\xa7\xfd\xcd\xcf\xa6\xee\x50\xd2\x81\xea\x30\x82\x28\x8c\xcf\xcc\x4b\xce\xcf\x2d\xc8\x49\x2d\x49\x45\xe8\x79\xb1\x67\xc7\xd3\x09\x1d\xcf\x36\x4f\x7d\x3e\xab\x05\x4d\x9b\x71\x7c\x72\x51\x6a\x62\x49\x6a\x7c\x5a\x62\x66\x0e\x48\xc7\xd3\xce\xad\xcf\x57\x77\x3f\x5d\xb2\xf1\xd9\xd4\xe9\x4a\x5c\xb5\x80\x00\x00\x00\xff\xff\x2f\x85\xf9\xb4\x91\x00\x00\x00")

func localesErrors_twJsonBytes() ([]byte, error) {
	return bindataRead(
		_localesErrors_twJson,
		"locales/errors_tw.json",
	)
}

func localesErrors_twJson() (*asset, error) {
	bytes, err := localesErrors_twJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/errors_tw.json", size: 145, mode: os.FileMode(420), modTime: time.Unix(1597372537, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesErrors_zhJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd2\xcb\x6e\xd3\x40\x14\x06\xe0\x7d\x9e\x22\xf2\xba\x8b\x34\xdc\xd9\x22\xb1\xe7\xb2\x3f\xb2\x9c\x09\x89\xda\xcc\x84\x19\x5b\x11\x42\x48\x34\x0d\xa9\x2b\xd2\x46\xa0\x84\x4b\x70\x95\x50\x11\x29\x6a\x45\x27\x2d\x95\x12\xdc\x04\x5e\xc6\x73\x3c\x5e\xf1\x0a\x68\x6a\x1a\x83\xc4\xd6\xfe\xbf\x7f\xce\xb1\xe7\x79\x2e\x9f\xb7\xd6\x0b\x85\xc2\x3a\x50\x06\x4f\x3d\xc2\x9f\x41\xdd\xe6\x76\xcd\xba\x9b\xb7\xf0\xec\x33\x06\xbb\xd1\x62\x94\xbc\x7c\xfb\xc0\xbc\x52\xdd\x26\xf6\xa7\xd6\xda\x15\x2a\xa6\x59\xa8\x52\x87\xd5\xea\x9b\xc4\x25\x86\xa5\x20\x1e\xb4\xd2\x78\x34\xdf\x53\x27\x1d\xec\x9f\x67\xee\x1a\x38\x9c\xd8\x2e\x81\xb2\x5d\xdd\x34\x44\xf9\x9f\xd4\x45\xa8\xbe\x9c\xea\xf3\x71\x16\xbb\xfe\x67\xa0\xab\x14\x0e\xc7\x5a\x1e\xa6\xa9\x5f\x8b\x8e\xda\x09\xe3\xfd\x69\x14\xb6\xcd\x09\x5d\x99\x1c\xb6\xe2\xe1\x38\x1e\xb4\x92\xde\x47\x2d\x65\x56\x73\x03\x1c\xc1\xcb\xe0\xb2\x0d\x42\xa1\x56\x15\x35\xdb\x75\x2a\xa6\xef\xde\xa3\x87\xf7\xf3\x8f\xcd\x63\xd3\xd0\xf9\x9e\xbc\xda\xcb\xd4\x4d\xe0\xc4\x61\xbc\x04\x94\xb9\x50\x66\x1e\x2d\x19\xa2\xe5\x0c\x4f\x9b\xf1\xa0\xa5\x4f\xa6\x6a\xd9\x37\xee\xeb\x07\x15\x4c\x32\x77\x0b\x6c\xc7\x21\x42\x80\x20\x42\x54\x19\xfd\xd7\xe3\xfb\x11\x7e\xeb\xeb\x1f\x13\xe5\x1f\x63\x73\x62\x8a\xc2\xe3\xe8\x62\x69\xd6\xe9\x4a\xbd\xbd\xc4\x7d\x1f\x0f\xb6\xd5\xec\x4c\xff\xdc\xc1\x60\x98\xf5\xde\x06\x8f\x6e\xb0\x06\x05\xc2\x39\xe3\x97\x1f\xad\xfb\x26\xee\x0d\xa3\xb0\x8d\xc1\xd1\xff\x37\xbf\xb3\x1a\x62\x85\xa2\xc5\x40\xcb\x03\xb5\x68\xaa\xf9\xdc\x5a\xcb\x99\x64\xf1\xf2\xf7\x7b\x82\x70\x60\x1c\xea\xb6\x10\x0d\xb3\x77\x83\x33\xfa\xc4\x90\xb8\x37\x41\x7f\x86\xfe\x3b\x25\xdb\xf1\x68\xeb\xef\x43\x8a\xe9\x25\xa8\x30\x4a\xc0\xf6\xdc\x0a\x38\xac\x44\x32\x89\xbb\xaf\x31\x08\x93\xa3\x8e\x96\x5b\x99\xcc\xbd\xc8\xfd\x0e\x00\x00\xff\xff\x21\xf7\x4d\x07\x77\x02\x00\x00")

func localesErrors_zhJsonBytes() ([]byte, error) {
	return bindataRead(
		_localesErrors_zhJson,
		"locales/errors_zh.json",
	)
}

func localesErrors_zhJson() (*asset, error) {
	bytes, err := localesErrors_zhJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/errors_zh.json", size: 631, mode: os.FileMode(420), modTime: time.Unix(1598717471, 0)}
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
	"conf/config.yaml":       confConfigYaml,
	"locales/error.go":       localesErrorGo,
	"locales/errors_en.json": localesErrors_enJson,
	"locales/errors_tw.json": localesErrors_twJson,
	"locales/errors_zh.json": localesErrors_zhJson,
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
	"conf": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{confConfigYaml, map[string]*bintree{}},
	}},
	"locales": &bintree{nil, map[string]*bintree{
		"error.go":       &bintree{localesErrorGo, map[string]*bintree{}},
		"errors_en.json": &bintree{localesErrors_enJson, map[string]*bintree{}},
		"errors_tw.json": &bintree{localesErrors_twJson, map[string]*bintree{}},
		"errors_zh.json": &bintree{localesErrors_zhJson, map[string]*bintree{}},
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
