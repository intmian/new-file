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

func textToNode(t string, nodes *Nodes) Node {
	// 因为加入serialnodes 需要外部引用 故加入一个指针
	switch t {
	case "date":
		return &DataNode{}
	case "time":
		return &TimeNode{}
	default:
		if t[0] == '@' {
			if t[1] == '@' {
				n, _ := strconv.Atoi(t[2:len(t)])
				return &SerialInputNode{no: n, outerValue: nodes.values}
			} else {
				// TODO
			}
		}
		return &InputNode{t}
	}

}
