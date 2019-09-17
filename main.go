/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/16 16:23
 */
package main

import "fileGenerate/tool"

func main() {
	str := tool.ReadFile(`template\hexo`)
	a := tool.AnalyserA{}
	a.SetText(str)
	nodes, _ := a.Out()
	print(nodes.Explain(&tool.Stdio{}))
}
