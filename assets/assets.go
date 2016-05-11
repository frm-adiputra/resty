// Code generated by go-bindata.
// sources:
// _assets/templates/build_handler.go.tmpl
// _assets/templates/declare_handler.tmpl
// _assets/templates/path_assign_handler.tmpl
// DO NOT EDIT!

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

var _templatesBuild_handlerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x94\x41\x6b\xdb\x30\x14\xc7\xcf\xd1\xa7\x78\x33\x3d\xd8\x2c\x55\xef\x83\x1e\x5c\x5b\x5d\x02\xa9\x1d\x1a\xe7\x34\x46\x10\xb5\x12\x0b\x52\xd9\xc8\xcf\x34\x43\xe8\xbb\x4f\x8a\x94\xae\x1b\xec\x98\x9e\xfc\x97\xde\x7b\xff\x9f\x9e\xc4\xb3\x31\xad\xd8\x4b\x25\x20\x29\x59\x55\xd4\xbb\xe7\x7a\xdb\xb0\x4d\x62\x2d\x31\x46\x73\x75\x10\x70\x83\xfc\x30\x87\x1b\xf8\x76\x0f\xb4\x1e\x84\xe6\x28\x7b\xf5\xf0\xab\xe1\x87\x5c\xb5\x6b\x8e\x9d\xfb\x3c\x09\xec\xfa\xd6\x15\xdd\xdd\x81\x31\xbe\xe2\xa3\xc1\xe0\x92\x2e\x0e\x70\x7b\x8e\xdc\xc2\x9b\xc4\x0e\xe8\x77\xd6\x58\xfb\x3a\x9d\xbc\x48\x07\xed\x8e\x72\xfa\x9a\xf8\x33\xa9\x97\xde\x7b\x87\x62\x6b\x93\x39\x60\x5f\xfa\xdd\x74\xc1\x55\x7b\x14\x7a\xa4\xc6\x4c\x83\x3b\x4f\xc1\x47\xf1\x28\xf5\x88\x2b\x81\x28\x34\xd0\x65\x69\x6d\x96\xcd\x1d\x45\xa8\xf6\x2f\xda\x7a\x1b\x69\x4e\x7c\x02\xad\xde\x5c\x70\x4e\x5d\x9f\x57\xb2\x15\x6b\x58\x20\x06\x7d\x7d\x66\xbd\x6e\x96\x75\xb5\x09\xd0\xb8\xb8\x3e\x75\xc1\xf2\x32\x20\xbd\xfa\x84\x97\xcc\x9b\x62\x11\x9f\xd2\xcb\x6b\x12\xff\xaf\x08\xf9\x33\xac\x8b\xbc\x2a\x57\xec\x79\xf7\xb8\xad\x8a\x5d\xc9\x8a\x55\x12\x86\x2f\x22\x41\x0b\x9c\xb4\x1a\x01\x3b\x01\x1d\xe2\x00\x5d\x0c\xec\x7b\xed\x36\xe5\x08\xf9\x7a\x49\xc9\x7e\x52\x2f\x97\x9a\xd8\x14\x8c\xa8\xa5\x3a\x64\x90\xfa\x3a\x1a\x83\x73\x10\x5a\xf7\x3a\x03\x43\x66\xee\x1a\xfc\x28\x9f\x3b\xa7\x95\x78\x7b\x9a\x4e\x69\x46\x66\xdd\x87\x4c\x1f\xf7\xb7\xf5\x30\xc9\x63\x9b\xfe\xf8\x19\x72\xa3\x97\x01\xdf\x13\x8a\xd7\xe1\xc8\xf1\xdf\x1f\x0f\x50\xd7\xc8\xcc\x3a\x3f\xb9\x3f\x3b\x7d\xb9\x07\x25\x8f\x9e\x3b\x0b\x4d\xf9\xe5\x19\xe2\xd2\x08\xb9\x6c\xbe\xd3\x5d\x94\xd8\xf7\x1b\xfd\x1d\x00\x00\xff\xff\x57\xad\xc5\x1e\xdf\x04\x00\x00")

func templatesBuild_handlerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesBuild_handlerGoTmpl,
		"templates/build_handler.go.tmpl",
	)
}

func templatesBuild_handlerGoTmpl() (*asset, error) {
	bytes, err := templatesBuild_handlerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/build_handler.go.tmpl", size: 1247, mode: os.FileMode(436), modTime: time.Unix(1462967674, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDeclare_handlerTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8f\x41\x6f\xfa\x30\x0c\xc5\xef\xfd\x14\x4f\x88\xe3\xff\x5f\xee\x93\x38\x30\x28\x02\xa9\xda\x26\xb4\xfb\xf0\x1a\xb7\x8d\xd4\xa6\x28\x31\x9b\x26\x2b\xdf\x7d\xe9\x28\x25\x27\x3b\x7e\x3f\xfb\x3d\x55\xc3\xb5\x75\x8c\xc5\xae\xd8\x96\x9b\x53\xf1\x71\xd8\xbc\xec\xca\xe2\xb4\x88\x31\x5b\xad\x70\x20\x67\x3a\xf6\x01\x5f\xe4\x2d\x7d\x76\x8c\x6a\x70\x42\xd6\x05\x50\xd7\xa1\xbd\x8d\x51\x0f\x1e\xd2\xda\x80\xcd\xdb\x31\x1f\xb9\xbd\xe5\xce\x04\x58\x27\xe9\x9f\x1f\x6b\x82\xbf\x56\x32\xca\x7b\x12\x90\x67\x34\xec\xd8\x93\xb0\x41\xed\x87\x1e\xe7\xe1\x32\xb6\x76\x70\x47\x73\xce\xb3\x74\xf5\xc1\xae\x11\xe4\x0f\xd7\x0c\xe9\xa9\x7a\x72\x0d\x63\x29\xd4\xfc\xc3\x12\x4f\x6b\xe4\xaf\x33\xbd\x7b\xfe\x79\xa7\x26\x85\x18\xa5\xc9\x90\xea\xa8\x9b\x7a\xd5\xff\x98\x60\x6b\xee\xec\x3c\xbb\x5e\xd2\x96\x2d\x05\xde\x5b\x1f\xa4\x64\x91\x94\x30\x09\x63\x44\xff\x5d\xb5\x29\x7c\x3e\x79\x52\x65\x67\x66\xee\x56\x47\x8d\xd9\xbd\xfe\x0d\x00\x00\xff\xff\x6d\x8a\x32\xd5\x5e\x01\x00\x00")

func templatesDeclare_handlerTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDeclare_handlerTmpl,
		"templates/declare_handler.tmpl",
	)
}

func templatesDeclare_handlerTmpl() (*asset, error) {
	bytes, err := templatesDeclare_handlerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/declare_handler.tmpl", size: 350, mode: os.FileMode(436), modTime: time.Unix(1462967564, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPath_assign_handlerTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x0a\x70\x0c\xf1\x88\x77\x0c\x0e\xf6\x74\xf7\x8b\xf7\x70\xf4\x73\xf1\x71\x0d\x52\xaa\xad\xe5\x52\x50\xa8\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd6\x51\x50\x29\x53\xb0\xb2\x55\xd0\xf3\x4d\x2d\xc9\xc8\x4f\x29\x06\x4b\x2b\x28\xe4\x96\x56\xe8\x55\x57\x97\x16\x14\xa4\x16\x39\x27\x16\x83\x14\xd6\xd6\x6a\x80\x0c\xce\x4b\xce\x0f\x48\x2c\xc9\x50\xd0\x03\x91\xb5\xb5\x3a\x40\xa3\xf4\xfc\x81\xaa\x12\x4b\x32\xf3\xf3\x3c\x5d\x6a\x6b\x35\xc1\xc6\xa7\xe6\xa5\x00\x4d\x82\xd1\x80\x00\x00\x00\xff\xff\x03\xe7\x6c\x71\x92\x00\x00\x00")

func templatesPath_assign_handlerTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPath_assign_handlerTmpl,
		"templates/path_assign_handler.tmpl",
	)
}

func templatesPath_assign_handlerTmpl() (*asset, error) {
	bytes, err := templatesPath_assign_handlerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/path_assign_handler.tmpl", size: 146, mode: os.FileMode(436), modTime: time.Unix(1462946519, 0)}
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
	"templates/build_handler.go.tmpl": templatesBuild_handlerGoTmpl,
	"templates/declare_handler.tmpl": templatesDeclare_handlerTmpl,
	"templates/path_assign_handler.tmpl": templatesPath_assign_handlerTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"build_handler.go.tmpl": &bintree{templatesBuild_handlerGoTmpl, map[string]*bintree{}},
		"declare_handler.tmpl": &bintree{templatesDeclare_handlerTmpl, map[string]*bintree{}},
		"path_assign_handler.tmpl": &bintree{templatesPath_assign_handlerTmpl, map[string]*bintree{}},
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

