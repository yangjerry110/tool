<!--
 * @Author: Jerry.Yang
 * @Date: 2023-04-27 11:35:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-20 10:45:52
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

## 假如需要使用protobuf，则需要提前安装protobuf相关
``` bash
    go install github.com/yangjerry110/tool/cmd/protoc-gen-go@latest
    go install github.com/yangjerry110/tool/cmd/protoc-gen-tool@latest
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
