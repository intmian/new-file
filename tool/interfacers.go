/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/17 14:01
 */
package tool

type IO interface {
	read() string
	write(s string)
}
type Node interface {
	explain(io IO) string // 将节点中的数据导出，如果需要数据就通过这种方式
}
type Analyzer interface {
	SetText(str string)
	canOut() bool
	Out() (Nodes, error)
}
