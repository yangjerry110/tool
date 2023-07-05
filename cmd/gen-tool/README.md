<!--
 * @Author: Jerry.Yang
 * @Date: 2023-04-27 11:35:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-23 15:00:51
 * @Description: 
-->

# cmd-tool
### gen-tool
    生成gen-tool 脚手架

## 如何使用
- 安装gen-tool 脚手架
``` bash
    go install github.com/yangjerry110/tool/cmd/gen-tool@latest
```

- 安装新应用
``` bash
    gen-tool new newApp
```
``` bash
    根据提示 input 项目地址
```
``` bash
    生成newApp之后，进入到newApp根目录，执行 `go mod tidy`
```

- 查看脚手架命令行
``` bash
    gen-tool --help
```
