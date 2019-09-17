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

func (a *AnalyserA) SetText(s string) {
	a.ori = s
}

func (a *AnalyserA) canOut() bool {
	return a.ori != ""
}

func (a *AnalyserA) Out() (Nodes, error) {
	if !a.canOut() {
		return Nodes{}, errors.New("no")
	}
	ptr := 0
	n := Nodes{}
	length := len(a.ori)
	for ptr != length {
		if ptr == length-2 {
			n.Add(textToNode(a.ori[ptr:length]))
		}

		c := a.ori[ptr]
		next := a.ori[ptr+1]
		switch c {
		case '<':
			if next == '<' {
				n.Add(&StrNode{"<"})
				ptr += 2
			} else {
				right := ptr + 1
				for a.ori[right] != '>' {
					right++
					if right == length {
						// TODO 语法错误
					}
				}
				n.Add(textToNode(a.ori[ptr+1 : right]))
				ptr = right + 1
			}
		case '>':
			if next == '>' {
				n.Add(&StrNode{">"})
				ptr += 2
			} else {
				// TODO 此处为模板语法错误
			}
		default: // 普通文字
			right := ptr
			for !(right == length-1) && (a.ori[right+1] != '<' || (a.ori[right] == '<' && a.ori[right+2] == '<')) {
				right++
			}
			n.Add(&StrNode{s: a.ori[ptr : right+1]})
			ptr = right + 1
		}

	}
	return n, nil
}
func textToNode(t string) Node {
	switch t {
	case "date":
		return &DataNode{}
	case "time":
		return &TimeNode{}
	default:
		return &InputNode{t}
	}

}
