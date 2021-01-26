package classpath

import "path/filepath"

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

/**
 * @Description: 使用-Xjre选项解析启动类路径和扩展类路径, 使用-classpath/-cp选项来解析用户类路径
 * @param jreOption
 * @param cpOption
 * @return *Classpath
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtraClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error)  {

}

func (self *Classpath) String() string {

}

func (self *Classpath) parseBootAndExtraClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath);
}

func (self *Classpath) parseUserClasspath(cpOption string) {

}