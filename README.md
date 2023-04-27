<!--
 * @Author: Jerry.Yang
 * @Date: 2022-09-19 17:46:05
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-27 11:31:14
 * @Description: 
-->
# tool
## 1.cmd
### 脚本文件
> cd 到cmd目录下查看相关说明

## 2.common
> 本意是此项目的一些公用方法，考虑到也可以public就也放到pkg中了


```
package main

import "github.com/yangjerry110/mytool/pkg/common"

func main() {
    accessToken,err := common.GetQiweiAccessToken(appId,cropId,cropSecret)
}
```

## 3.conf
>1.对一些配置文件进行解析，需要提前定义到需要解析的文件和解析对应的结构体  
>2.配置文件定义(以yaml文件示例)

```
conf_path: "conf_path"
```

引用实例：

```
package main

import "github.com/yangjerry110/mytool/pkg/conf"

func main() {

    type MyConf struct {
        ConfPath string `yaml:"conf_path"`
    }


    // 解析yaml的配置到MyConf结构体
    myConf := &MyConf{}
    yamlConf.GetYamlConf(yamlConfPath,myConf)
    myConfPath := myConf.ConfPath

}
```

## 4.http
> 1.http调用相关的  
> 2.需要提前定义好input, output   
> 3.ps : options 配置需要引导 http.Option下   

```
package main 

import (
    httpPkg "github.com/yangjerry110/mytool/pkg/http" 
    "github.com/yangjerry110/mytool/http"
)

func main() {

    // 以json数据示例
    type Input struct {
        InputStr string
    }

    type OutPut struct {
        RetCode int 
        RetMsg string
        RetResult bool
    }

    jsonReq,_ := json.Marshal(Input{InputStr:"inputStr"})

    // 定义options
    httpOptions := http.HttpOptions{}
    options := []http.HttpOptionFunc{
        httpOptions.SetHeaders(map[string]string{
            "Content-Type": "multipart/form-data",
        })
    }

    resp := &OutPut{}
    httpPkg.HttpRequest(method,url,bytes.NewBuffer(jsonReq),resp,options)
}

```

## 5.extenrnal

> 1.qiwei相关   
> 2.以企微通知示例   
> 3.支持 text,markdown,image,news,card    
> 4.bot通知 支持 text,markdown,image,news

```
package main 

import "github.com/yangjerry110/mytool/pkg/extenrnal"

func main() {

    /**
    * @param {string} AppId appId
    * @param {string} MsgType 消息类型
    * @param {string} CropId 公司id
    * @param {string} CropSecret 公司秘钥
    * @param {string} AgentId 通知的qiwei应用的id
    * @param {string} DepartmentIds 通知的组织架构集合
    * @param {string} TagIds 通知的Tag集合
    * @param {string} UserIds 通知的人
    * @param {int32} Safe safe
    * @param {string} SendMsg 通知的消息
    * @param {string} MediaData 通知的媒体消息内容
    * @param {string} MediaType 通知的媒体消息类型
    * @param {string} Title 标题
    * @param {string} Description 简介
    * @param {string} Url 链接
    * @param {string} PicUrl 图片链接
    * @param {int32} EnableIdTrans
    * @param {string} Btntxt news消息时候的按钮
    * @param {string} AppletId AppletId 小程序id
    * @param {string} AppletPagepath AppletPagepath 小程序链接
    * @param {string} QiweiFilePath 通知媒体消息的时候，存放媒体内容的地址
     */
    qiweiNotice := extenrnal.qiwei.QiweiNoticePkg{}

    result,err := qiweiNotice.QiweiNotice()

}
```

## 6.perm 

> 1.加密解密相关    
> 2.PS: 使用rsa加解密之前，需要先生成公钥和私钥

```

package main 

import "github.com/yangjerry110/mytool/pkg/perm"

func main() {
    encrtyStr,err := perm.RsaDecrty(permPath,inputStr)
}

```

## 7.logger
> 1.默认采用的logrus    
> 2.采用的是插拔式，可以设置option，默认是json格式的输出式的日志

```
package main 

import （
    "github.com/yangjerry110/mytool/pkg/logger"
    toolLogger "github.com/yangjerry110/mytool/logger"
）

func main() {

    // 默认是logrus
    // 可以自己设置logger引擎
    logger.setLogger()

    // 可以自己设置option
    // 但是调用对应的log方法的时候，level会被重置掉
    logger.SetOptions([]toolLogger.LoggerOptionFunc{
		logger.SetIsReportcaller(true),
		logger.SetLevel(logger.Level(toolLogger.DebugLevel)),
	}).WithField("testFields", "testFieldVal").Info("this is test withField")

    // show this
    // {"file":"/Users/admin/go/src/my-tool/logger/logrusLog.go:60","func":"github.com/yangjerry110/mytool/logger.(*LogrusLog).WriteLog","level":"info","msg":"this is test withField","time":"2022-10-09T17:32:59+08:00"}

}
```

## 8.cache
>1.默认采用的redis   
>2.可以自己设置cache的引擎(setCache func)   
>3.可以单独获取结果和error(getResult func, getErr func) 
>4.使用默认的方式连接引擎的话，需要使用默认的配置(有相应的default)  
>5.连接cache引擎的时候，需要传入conf(以yaml的配置方式)


```
package main 
import (
    "github.com/yangjerry110/mytool/pkg/cache"
)

func main() {
    cache.Client(redisConfPath).Set("test","test",1*time.Minute).GetErr()

    resultInterface,err := cache.Client(redisConfPath).Get("test").Result()
}
```






