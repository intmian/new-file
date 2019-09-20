/*
* @Describe: 为此程序写的reg生成器
* @Author:   mian
* @Date:     2019/9/20 21:09
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reT := `Windows Registry Editor Version 5.00  
[HKEY_CLASSES_ROOT\Directory\Background\shell\newFile]  
@="调用模板生成器"
[HKEY_CLASSES_ROOT\Directory\Background\shell\newFile\command]  
@="%s"`
	unre := `Windows Registry Editor Version 5.00
[-HKEY_CLASSES_ROOT\Directory\Background\shell\newFile]-
[-HKEY_CLASSES_ROOT\Directory\Background\shell\newFile\command]-`
	exePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	exePath += `\main.exe`
	exePath = strings.ReplaceAll(exePath, `\`, `\\`) // 在reg脚本运行时还要一层转义
	re := fmt.Sprintf(reT, exePath)
	_ = ioutil.WriteFile("a注册.reg", []byte(re), 0666)
	_ = ioutil.WriteFile("a解除注册.reg", []byte(unre), 0666)
}
