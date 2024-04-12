<!--
 * @Author: Jerry.Yang
 * @Date: 2023-04-27 11:35:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 17:04:22
 * @Description: 
-->

## cmd-tool
### gen-tool
    生成gen-tool 脚手架
### protoc-gen-tool
    生成protoc-gen-tool 插件

## 如何使用
- 安装gin-framwwork 脚手架
``` bash
    go install github.com/yangjerry110/tool/cmd/protoc-gen-tool@latest
    go install github.com/yangjerry110/too/cmd/gen-tool@latest
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
``` base
    获取命令行 gen-tool --help
```