/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/17 17:25
 */
package tool

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) string {
	f, _ := os.Open(path)
	b, _ := ioutil.ReadAll(f)
	return string(b)
}

func FillTemp(temp string, a Analyzer, io IO) string {
	// 因为Analyzer其实对应的是指针形的，所以如此写
	a.SetText(temp)
	n, _ := a.Out()
	return n.Explain(io)
}

func SimpleFillTemp(tempName string) string {
	s := ReadFile(tempName)
	return FillTemp(s, &AnalyserA{}, &StdIOSingle)
}

func (n *Nodes) textToNode(t string) Node {
	// 因为加入serialNodes 需要外部引用 故加入一个接收者
	switch t {
	case "date":
		return &DataNode{}
	case "time":
		return &TimeNode{}
	default:
		length := len(t)
		if t[0] == '@' { // serialNode
			if t[1] == '@' {
				no, _ := strconv.Atoi(t[2:length])
				return &SerialInputNode{
					no: no,
					outerValue: n.values,
					ifPrimary: false}
			} else { // primary node
				args := strings.Split(t[1:length], " ") // TODO
				no, err1 := strconv.Atoi(args[0])
				if err1 != nil {
					// TODO 模板语法错误
				}
				return &SerialInputNode{
					no:         no,
					str:        args[1],
					ifPrimary:  true,
					outerValue: n.values}
			}
		} else {
			switch t[0:4] {
			case "file":
				return &FileNode{path: t[5:length]}
			}
		}
		return &InputNode{t}
	}

}
