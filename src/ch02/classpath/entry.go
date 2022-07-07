package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

// 表示类路径选项
type Entry interface {
	// 负责寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)
	// 相当于Java中的toString()方法
	String() string
}

// 根据参数创建不同类型的entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
