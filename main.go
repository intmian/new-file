/*
* @Describe:
* @Author:   mian
* @Date:     2019/9/16 16:23
 */
package main

import (
	"fileGenerate/tool"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	/*
		相当地址指的是命令行调用的地址
		只有通过参数第一项获得的参数才是有用的
	*/
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	tempFolder := dir + `\template`
	files, _ := ioutil.ReadDir(tempFolder)
	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}
	println("请选择模式1.低级 2.高级")
	var mode int
	_, _ = fmt.Scan(&mode)
	switch mode {
	case 1:
		println("请选择模板")
		for i, name := range fileNames {
			fmt.Printf("%d %s\n", i+1, name)
		}
		var no int
		_, _ = fmt.Scan(&no)
		no--
		println("输入文件名")
		var s string
		_, _ = fmt.Scan(&s)
		_ = ioutil.WriteFile(s, []byte(tool.SimpleFillTemp(dir+`\template\`+fileNames[no])), 0666)
		exec.Command(s)
	default:
		println("无此模式")
	}
}
