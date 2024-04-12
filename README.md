<!--
 * @Author: Jerry.Yang
 * @Date: 2023-11-30 16:16:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 16:49:16
 * @Description: 
-->
## tool

### 1.cmd
> 查看文档[https://github.com/yangjerry110/tool/tree/master/cmd]

### 2.cache
>1.设置redis的conf
>2.创建cache实例化

代码示例
```
package main
import "github.com/yangjerry110/tool/cache"

func main() {

   // 设置redis的conf
   err := cache.CreateRedisConf().SetConfig()
   // 创建redis的实例化
   cacheInterface := cache.CreateRedisCache()
}
```

### 3.db
>1.设置db的conf
>2.创建db的实例化

代码示例
```
package main
import "github.com/yangjerry110/tool/db"

func main() {

    // 设置db的conf
    err := db.CreateGormDbConf().SetConfig()
    // 创建Gorm的实例化
    dbInterface := db.CreateGormDb()
}
```

### 4.logger
>1.设置logrus logger option的conf
>2.创建logger的实例化
>3.默认使用了logrus的引擎，可以开箱即用
```
package main
import "github.com/yangjerry110/tool/logger"

func main() {
    // 设置logrus logger option的conf
    err := logger.CreateLogrusOptionConf()

    // 设置logger的引擎，实例化
    logger.SetLoggerEnginee(LoggerInterface)

    // info
    logger.Info("this is info log")
    logger.Infof("this is %s log","info")
}
```

### 5.perm
>1.创建perm的实例化

```
package main
import "github.com/yangjerry110/tool/perm"

func main() {

    // 设置perm的实例化
    permInterface := perm.CreateRsaPerm()
}
```

### 6.router
>1.设置gin的conf(内置了redis的conf，gormDb的conf，还有gin的conf)
>2.创建router的实例化
>3.注册router

```
package main
import "github.com/yangjerry110/tool/router"

func main() {

    // 设置gin的conf
    err := router.CreateGinConf().SetConfig()

    // 创建gin的实例化
    routerInterface := router.CreateGinRouter()

    // 注册router
    routerInterface.Register("routerName",routerRegisterInterface)
}
```





