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
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/yangjerry110/tool/toolerrors"
	"gopkg.in/yaml.v3"
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
	mu      sync.Mutex
	isWatch bool              // 是否已启动监听
	watcher *fsnotify.Watcher // 全局共享的 watcher 实例
}

// watchConf 全局变量，存储监听状态
var watchConf = &watchConfing{}

// watchFilesConf 全局变量，存储所有被监听的文件信息（使用 sync.Map 实现线程安全）
var watchFilesConf = sync.Map{}

// watchDirsConf 全局变量，存储已经加入 watcher 的配置目录
var watchDirsConf = sync.Map{}

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
	configFile, err := filepath.Abs(filepath.Join(w.watchfile.filePath, w.watchfile.fileName))
	if err != nil {
		return err
	}
	configFileMd5, err := w.getConfigFileMd5(configFile)
	if err != nil {
		return err
	}
	configDir := filepath.Dir(configFile)

	watchFileObj := *w.watchfile
	watchFileObj.filePath = configDir
	watchFileObj.fileName = filepath.Base(configFile)

	// 将文件信息存储到全局变量中
	watchFilesConf.Store(configFileMd5, &watchFileObj)

	watchConf.mu.Lock()
	defer watchConf.mu.Unlock()

	if watchConf.isWatch {
		// watcher 已启动，直接将新目录追加到现有 watcher
		if err := w.addWatchDir(watchConf.watcher, configDir); err != nil {
			fmt.Printf("watchConf AddWatch Err : %+v; configDir : %+v\r\n", err, configDir)
			return err
		}
		return nil
	}

	// 初始化 watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	watchConf.watcher = watcher

	// 将当前配置目录加入监听。配置中心常用 rename 覆盖文件，监听文件本身会在 inode 替换后失效。
	if err := w.addWatchDir(watcher, configDir); err != nil {
		watcher.Close()
		return err
	}

	// 标记已启动监听
	watchConf.isWatch = true

	// 启动文件监听协程
	go w.watchFileLoop(watcher)
	return nil
}

/**
 * @description: watchFileLoop 监听文件变化，并在文件修改时重新加载配置
 * @receiver w *watch 监听对象
 * @param watcher *fsnotify.Watcher 共享的 fsnotify 监听器
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:46:23
 */
func (w *watch) watchFileLoop(watcher *fsnotify.Watcher) {
	defer watcher.Close()

	for {
		select {
		case event, ok := <-watcher.Events:
			// 如果事件通道关闭，直接返回
			if !ok {
				return
			}

			// Write：直接写入；Create/Rename：配置中心或编辑器先写临时文件再替换目标文件。
			if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) || event.Has(fsnotify.Rename) || event.Has(fsnotify.Remove) {
				for _, watchFileObj := range w.getEventWatchFiles(event.Name) {
					configFile := filepath.Join(watchFileObj.filePath, watchFileObj.fileName)
					if err := w.reloadConfigFile(configFile, watchFileObj); err != nil {
						fmt.Printf("watchConf Reload Err : %+v; configFile : %+v\r\n", err, configFile)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			// 处理监听器错误
			if !ok {
				return
			}
			fmt.Printf("watchConf Err : %+v;\r\n", err)
		}
	}
}

/**
 * @description: addWatchDir 将配置目录加入 watcher，避免重复监听同一个目录
 * @receiver w *watch 监听对象
 * @param watcher *fsnotify.Watcher 共享的 fsnotify 监听器
 * @param configDir string 配置目录
 * @return error 如果目录监听失败，返回错误
 * @author: Jerry.Yang
 * @date: 2026-04-30 00:00:00
 */
func (w *watch) addWatchDir(watcher *fsnotify.Watcher, configDir string) error {
	if _, ok := watchDirsConf.Load(configDir); ok {
		return nil
	}

	if err := watcher.Add(configDir); err != nil {
		return err
	}
	watchDirsConf.Store(configDir, struct{}{})
	return nil
}

/**
 * @description: getEventWatchFiles 获取事件影响的配置文件
 * @receiver w *watch 监听对象
 * @param eventName string fsnotify 事件文件名
 * @return []*watchFile 需要重新加载的配置文件列表
 * @author: Jerry.Yang
 * @date: 2026-04-30 00:00:00
 */
func (w *watch) getEventWatchFiles(eventName string) []*watchFile {
	configFile, err := filepath.Abs(eventName)
	if err != nil {
		configFile = filepath.Clean(eventName)
	}

	configFileMd5, err := w.getConfigFileMd5(configFile)
	if err == nil {
		if watchfile, isExistWatchFile := watchFilesConf.Load(configFileMd5); isExistWatchFile {
			return []*watchFile{watchfile.(*watchFile)}
		}
	}

	// 配置中心可能只上报临时文件 rename 事件，这里对同目录配置做兜底 reload。
	configDir := filepath.Dir(configFile)
	watchFiles := make([]*watchFile, 0)
	watchFilesConf.Range(func(_, value interface{}) bool {
		watchFileObj := value.(*watchFile)
		if watchFileObj.filePath == configDir {
			watchFiles = append(watchFiles, watchFileObj)
		}
		return true
	})
	return watchFiles
}

/**
 * @description: reloadConfigFile 重新加载配置文件内容
 * @receiver w *watch 监听对象
 * @param configFile string 配置文件路径
 * @param watchFileObj *watchFile 被监听文件信息
 * @return error 如果读取或解析失败，返回错误
 * @author: Jerry.Yang
 * @date: 2026-04-30 00:00:00
 */
func (w *watch) reloadConfigFile(configFile string, watchFileObj *watchFile) error {
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	fmt.Printf("Config file %s modified. Reloading...\r\n", configFile)

	// 重新加载配置文件（只调用 yaml 解析，避免再次注册 watcher）
	fileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(fileContent, watchFileObj.confData); err != nil {
		return err
	}
	return nil
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
