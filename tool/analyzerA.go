/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/17 14:33
 */
package tool

import "errors"

type AnalyserA struct {
	ori   string
	isPro bool
}

func (a *AnalyserA) GetText(s string) {
	a.ori = s
}

func (a *AnalyserA) canOut() bool {
	return a.ori != ""
}

func (a *AnalyserA) out() (Nodes, error) {
	if !a.canOut() {
		return Nodes{}, errors.New("no")
	}
	ptr := 0
	n := Nodes{}
	for ptr != len(a.ori) {
		c := a.ori[ptr]
		next := a.ori[ptr+1]
		switch c {
		case '<':
			if next == '<' {
				ptr += 2

			}
		case '>':
		default: // 普通文字
			right := ptr
			for a.ori[right+1] != '<' || (a.ori[right] == '<' && a.ori[right+2] == '<') {
				right++
			}
			ptr = right + 1
		}

	}
}
func textToNode(t string) Node {
	switch t {
	case "date":
		return &DataNode{}
	case "time":
		return &TimeNode{}
	}
}
