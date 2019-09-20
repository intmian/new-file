/*
* @Describe: 默认的分析器
* @Author:   mian
* @Date:     2019/9/17 14:33
 */
package tool

import "errors"

type AnalyserA struct {
	ori   string
	isPro bool
}

// 设置分析的源文本
func (a *AnalyserA) SetText(s string) {
	a.ori = s
}

func (a *AnalyserA) canOut() bool {
	return a.ori != ""
}

// 分析过程
func (a *AnalyserA) Out() (Nodes, error) {
	if !a.canOut() {
		return Nodes{}, errors.New("no")
	}
	ptr := 0
	n := NewNodes()
	length := len(a.ori)
	for ptr != length {
		c := a.ori[ptr]
		if ptr == length-2 {
			switch c {
			case '<', '>':
				n.Add(n.textToNode(string(c)))
			default:
				n.Add(n.textToNode(a.ori[ptr:length]))
			}
			break
		}

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
				n.Add(n.textToNode(a.ori[ptr+1 : right]))
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
			for !(right == length-1) && a.ori[right+1] != '<' && a.ori[right+1] != '>' {
				right++
			}
			n.Add(&StrNode{s: a.ori[ptr : right+1]})
			ptr = right + 1
		}

	}
	return *n, nil
}
