<!--
 * @Author: Jerry.Yang
 * @Date: 2023-04-27 11:35:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-20 11:18:16
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

- 假如需要使用protobuf，则需要提前安装protobuf相关
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

## 命令行
- newApi
```

//第一步 定义protobuf文件(PS 必须定义在protobuf文件夹下，因为生成的时候会指定protobuf文件夹，有继承关系)
//第二步 根据提示回答，假如是firstCreate，会在base文件中添加相关的CreateFunc
//第三步 假如是append，不会重新生成service，会扫描没生成的func，进行添加
//第四步 根据生成的文件自己修改
gen-tool newApi demo.proto

```
