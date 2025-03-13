/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 14:11:24
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 22:28:23
 * @Description: 配置文件监听模块，提供配置文件的热更新功能。
 * 支持通过 fsnotify 监听配置文件的修改事件，并在文件修改时自动重新加载配置。
 */
package conf

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/yangjerry110/tool/toolerrors"
)

// watch 结构体，用于管理配置文件监听功能
type watch struct {
	watchfile *watchFile // 被监听的文件信息
}

// watchFile 结构体，存储被监听文件的详细信息
type watchFile struct {
	filePath string      // 文件路径
	fileName string      // 文件名称
	fileType string      // 文件类型（如 yaml）
	confData interface{} // 配置文件的数据结构
}

// watchConfing 结构体，用于管理监听状态
type watchConfing struct {
	isWatch bool // 是否已启动监听
}

// watchConf 全局变量，存储监听状态
var watchConf = &watchConfing{}

// watchFilesConf 全局变量，存储所有被监听的文件信息（使用 sync.Map 实现线程安全）
var watchFilesConf = sync.Map{}

/**
 * @description: SetConfig 实现 Conf 接口，用于启动配置文件监听
 * @receiver w *watch 监听对象
 * @return error 如果监听文件为空，返回错误；否则启动监听并返回 nil
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:48:02
 */
func (w *watch) SetConfig() error {
	// 检查监听文件是否为空
	if w.watchfile == nil {
		return toolerrors.New("conf Err : watch conf no watchFile")
	}

	// 计算文件的 MD5 值，作为唯一标识
	configFile := fmt.Sprintf("%s/%s", w.watchfile.filePath, w.watchfile.fileName)
	configFileMd5, err := w.getConfigFileMd5(configFile)
	if err != nil {
		return err
	}

	// 将文件信息存储到全局变量中
	watchFilesConf.Store(configFileMd5, w.watchfile)

	// 如果已启动监听，直接返回
	if watchConf.isWatch {
		return nil
	}

	// 启动文件监听协程
	go w.watchFile()

	// 标记已启动监听
	watchConf.isWatch = true
	return nil
}

/**
 * @description: watchFile 监听文件变化，并在文件修改时重新加载配置
 * @receiver w *watch 监听对象
 * @return error 如果监听器初始化失败，返回错误；否则持续监听文件变化
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:46:23
 */
func (w *watch) watchFile() error {
	// 初始化 fsnotify 监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("watchConf watchFile Err : %+v;", err)
		fmt.Print("\r\n")
		return err
	}
	defer watcher.Close()

	// 遍历所有被监听的文件，添加到监听器中
	watchFilesConf.Range(func(watchFileName, watchfile any) bool {
		watchFileObj := watchfile.(*watchFile)
		err := watcher.Add(fmt.Sprintf("%s/%s", watchFileObj.filePath, watchFileObj.fileName))
		if err != nil {
			fmt.Printf("watchConf watchFile AddWatch Err : %+v;", err)
			fmt.Print("\r\n")
			return false
		}
		return true
	})

	// 处理文件变化事件
	for {
		select {
		case event, ok := <-watcher.Events:
			// 如果事件通道关闭，直接返回
			if !ok {
				return nil
			}

			// 如果文件被修改
			if event.Has(fsnotify.Write) {
				fmt.Printf("Config file %s modified. Reloading...\n", event.Name)
				fmt.Print("\r\n")
				configFile := event.Name

				// 获取文件扩展名
				extension := strings.ToLower(filepath.Ext(configFile))

				// 如果是 YAML 文件
				if extension == ".yaml" {
					// 计算文件的 MD5 值
					configFileMd5, err := w.getConfigFileMd5(configFile)
					if err != nil {
						return err
					}

					// 获取文件信息
					watchfile, isExistWatchFile := watchFilesConf.Load(configFileMd5)
					if !isExistWatchFile {
						fmt.Printf("watchFile is not exist; configFile : %+v", configFile)
						fmt.Print("\r\n")
						return toolerrors.New("conf Err : watch conf no confFile")
					}

					// 重新加载配置文件
					watchFileObj := watchfile.(*watchFile)
					if err := CreateConf(&yamlConf{
						filePath: watchFileObj.filePath,
						fileName: watchFileObj.fileName,
						fileType: watchFileObj.fileType,
						confData: &watchFileObj.confData,
					}).SetConfig(); err != nil {
						fmt.Printf("watchConf SetYamlConf Err : %+v; configFile : %+v", err, configFile)
						fmt.Print("\r\n")
						// 即使加载失败，也不中断监听
					}
				}
			}
		case err, ok := <-watcher.Errors:
			// 处理监听器错误
			fmt.Printf("watchConf Err : %+v;", err)
			fmt.Print("\r\n")
			if ok {
				return nil
			}
		}
	}
}

/**
 * @description: getConfigFileMd5 计算文件的 MD5 值
 * @param {string} configFile 文件路径
 * @return {string} 文件的 MD5 值
 * @return {error} 如果计算失败，返回错误
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:39:03
 */
func (w *watch) getConfigFileMd5(configFile string) (string, error) {
	newMd5 := md5.New()
	newMd5.Write([]byte(configFile))
	hashInBytes := newMd5.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}
