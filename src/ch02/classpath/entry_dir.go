package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	// 存放绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	// 先把参数转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 先把目录和class文件名拼成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
