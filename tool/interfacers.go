/*
* @Describe: 一些接口
* @Author:   mian
* @Date:     2019/9/17 14:01
 */
package tool


// 节点交互接口
type IO interface {
	read() string
	write(s string)
}

//节点
type Node interface {
	explain(io IO) string // 将节点中的数据导出，如果需要数据就通过这种方式
}

// 将文本解析为节点集的分析器
type Analyzer interface {
	SetText(str string)
	canOut() bool
	Out() (Nodes, error)
}
