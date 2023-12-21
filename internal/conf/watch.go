/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 14:11:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 14:46:26
 * @Description: watch
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
	"github.com/yangjerry110/tool/internal/errors"
)

type Watch struct {
	WatchFile *WatchFile
}

type WatchFile struct {
	FilePath string
	FileName string
	FileType string
	ConfData interface{}
}

type watchConf struct {
	IsWatch bool
}

/**
 * @description: WatchConf
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:17:24
 * @return {*}
 */
var WatchConf = &watchConf{}

/**
 * @description: WatchFilesConf
 * @author: Jerry.Yang
 * @date: 2023-12-20 16:49:49
 * @return {*}
 */
var WatchFilesConf = sync.Map{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:48:02
 * @return {*}
 */
func (w *Watch) SetConfig() error {

	// If w.WatchFile == nil
	// Return err
	if w.WatchFile == nil {
		return errors.ErrWatchConfNoWatchFile
	}

	// Add WatchFile
	// md5 filePath + fileName
	configFile := fmt.Sprintf("%s/%s", w.WatchFile.FilePath, w.WatchFile.FileName)

	// Get configFile Md5
	configFileMd5, err := w.getConfigFileMd5(configFile)
	if err != nil {
		return err
	}

	// Set WatchFilesConf
	WatchFilesConf.Store(configFileMd5, w.WatchFile)

	// If IsWatch
	// If IsWatch ; return
	if WatchConf.IsWatch {
		return nil
	}

	// Go watchFile
	go w.watchFile()

	// The first
	// Set IsWatch == true
	WatchConf.IsWatch = true
	return nil
}

/**
 * @description: watchFile
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:46:23
 * @return {*}
 */
func (w *Watch) watchFile() error {

	// NewWatcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("watchConf watchFile Err : %+v;", err)
		fmt.Print("\r\n")
		return err
	}
	defer watcher.Close()

	// Monitor config files
	WatchFilesConf.Range(func(watchFileName, watchFile any) bool {

		// Set interface
		watchFileObj := watchFile.(*WatchFile)
		err := watcher.Add(fmt.Sprintf("%s/%s", watchFileObj.FilePath, watchFileObj.FileName))
		if err != nil {
			fmt.Printf("watchConf watchFile AddWatch Err : %+v;", err)
			fmt.Print("\r\n")
			return false
		}
		return true
	})

	// Handle file change events
	for {
		select {
		case event, ok := <-watcher.Events:
			// Get Watch.Events
			// If !ok
			if !ok {
				return nil
			}

			// If fsnotify.Write
			if event.Has(fsnotify.Write) {
				fmt.Printf("Config file %s modified. Reloading...\n", event.Name)
				fmt.Print("\r\n")
				configFile := event.Name

				// Get configFile extension
				extension := strings.ToLower(filepath.Ext(configFile))

				// If .yaml
				if extension == ".yaml" {

					// Get configFileMd5
					configFileMd5, err := w.getConfigFileMd5(configFile)
					if err != nil {
						return err
					}

					// Get WatchFile
					watchFile, isExistWatchFile := WatchFilesConf.Load(configFileMd5)
					if !isExistWatchFile {
						fmt.Printf("watchFile is not exist; configFile : %+v", configFile)
						fmt.Print("\r\n")
						return errors.ErrWatchConfNoConfFile
					}

					// Set Conf
					watchFileObj := watchFile.(*WatchFile)
					if err := CreateConf(&Yaml{FilePath: watchFileObj.FilePath, FileName: watchFileObj.FileName, FileType: watchFileObj.FileType, ConfData: watchFileObj.ConfData}).SetConfig(); err != nil {
						fmt.Printf("watchConf SetYamlConf Err : %+v; configFile : %+v", err, configFile)
						fmt.Print("\r\n")
						// return err
					}

				}
			}
		// Have err
		case err, ok := <-watcher.Errors:
			fmt.Printf("watchConf Err : %+v;", err)
			fmt.Print("\r\n")
			if ok {
				return nil
			}
		}
	}
}

/**
 * @description: getConfigFileMd5
 * @param {string} configFile
 * @author: Jerry.Yang
 * @date: 2023-12-20 14:39:03
 * @return {*}
 */
func (w *Watch) getConfigFileMd5(configFile string) (string, error) {
	newMd5 := md5.New()
	newMd5.Write([]byte(configFile))
	hashInBytes := newMd5.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}
