/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/17 17:27
 */
package tool

import "fmt"

var StdIOSingle Stdio

func init() {
	StdIOSingle = Stdio{}
}

type Stdio struct {
}

func (s *Stdio) read() string {
	var str string
	_, _ = fmt.Scan(&str)
	return str
}

func (s *Stdio) write(str string) {
	println(str)
}
