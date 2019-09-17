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
	"path/filepath"
)

func main() {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	myFolder := `template`
	files, _ := ioutil.ReadDir(myFolder)
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
			fmt.Printf("%d %s\n", i, name)
		}
		var no int
		_, _ = fmt.Scan(&no)
		println("输入文件名")
		var s string
		_, _ = fmt.Scan(&s)
		_ = ioutil.WriteFile(dir+`\`+s, []byte(tool.SimpleFillTemp(fileNames[no])), 0666) //写入文件(字节数组)
	default:
		println("无此模式")
	}
}
