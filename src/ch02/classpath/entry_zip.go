package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
)
import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开压缩文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	// 遍历压缩包内的文件
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
