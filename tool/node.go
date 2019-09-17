/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/17 14:05
 */
package tool

import (
	"fmt"
	"time"
)

type Nodes struct {
	nodes []Node
}

func (n *Nodes) Add(node Node) {
	n.nodes = append(n.nodes, node)
}
func (n *Nodes) Explain(io IO) string {
	s := ""
	for _, n := range n.nodes {
		s += n.explain(io)
	}
	return s
}

type DataNode struct {
}

func (d *DataNode) explain(io IO) string {
	t := time.Now()
	return fmt.Sprintf("%4d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

type TimeNode struct {
}

func (t *TimeNode) explain(io IO) string {
	n := time.Now()
	return fmt.Sprintf("%2d:%2d:%2d", n.Hour(), n.Minute(), n.Second())
}

type InputNode struct {
	s string
}

func (i *InputNode) explain(io IO) string {
	io.write("请输入 " + i.s)
	return io.read()
}

type StrNode struct {
	s string
}

func (s *StrNode) explain(io IO) string {
	return s.s
}
