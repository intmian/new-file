/*
* @Describe: 为此程序写的reg生成器
* @Author:   mian
* @Date:     2019/9/20 21:09
 */
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var bom []byte

func init() {
	bom = []byte{byte(0xEF), byte(0xBB), byte(0xBF)}
}
func AddBom(b []byte) []byte {
	return bytes.Join([][]byte{bom, b}, []byte(""))
}
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

	_ = ioutil.WriteFile("注册.reg", AddBom([]byte(re)), 0666)
	_ = ioutil.WriteFile("解除注册.reg", AddBom([]byte(unre)), 0666)
}
