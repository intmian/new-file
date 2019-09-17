/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/17 17:25
 */
package tool

import (
	"io/ioutil"
	"os"
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
	s := ReadFile(`template\` + tempName)
	return FillTemp(s, &AnalyserA{}, &StdIOSingle)
}
