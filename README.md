# new-file

这是一个用来自动生成文档的命令行小工具，由于规模较小加自用，readme就随便写了。

## 如何用

1. 首先下载[压缩包](https://github.com/intmian/new-file/releases/download/1.1/new_file.rar)
2. 解压后双击register.exe,会产生两个reg文件
3. 点击注册即可将程序加入右键菜单。
4. 当你想要删除程序时，点两下解除注册就可以删除右键菜单中这项。
5. 在空白处用右键点击背景选择调用模板生成器即可使用。

> 改变程序路径时，应先点击解除注册再从上述第二步开始执行。

## 模板语法

模板语法与汉语一致。

此外

| 文法           | 作用                                                         | 示例         |
| -------------- | ------------------------------------------------------------ | ------------ |
| \<date>         | 将用当前日期填充                                             | \            |
| \<time>         | 将用当前时间填充                                             | \            |
| <(str)>        | 在生成文档时将要求输入，并填充                | \<input this> |
| <@(int) (str)> | 在生成文档时将要求输入，并将此文本以编号为key保存，并填充 | <@1 input>   |
| <@@(int)>       | 以编号为key取出值，并填充                           | <@@1>         |
| <file (path)>  | 将命令行当前目录下的文件读入，并填充                | <file 1.py>  |
| <<             | 转义为<                                                      | \            |
| >>             | 转义为>                                                      | \            |

将你写好的模板文件放在执行文件同层的template文件夹下即可
