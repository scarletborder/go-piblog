package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"piblog/src/model"
	"piblog/src/upload"
	htmlparser "piblog/src/upload/parser/html_parser"
)

func HandleMigrate(dirPath string, tag string) error {
	// 检查路径是否存在以及是否是目录
	fileInfo, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("路径不存在: %s", dirPath)
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("路径不是一个文件夹: %s", dirPath)
	}

	// 读取文件夹中的所有文件
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("无法读取文件夹内容: %v", err)
	}

	// 遍历文件夹中的文件
	for _, file := range files {
		// 只处理后缀为 .html 的文件
		if filepath.Ext(file.Name()) == ".html" {
			filePath := filepath.Join(dirPath, file.Name())

			// 读取文件内容
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("无法读取文件 %s: %v\n", filePath, err)
				continue
			}

			// 将文件内容转换为字符串，并通过 HandleString 处理
			d := htmlparser.ParseDocument(string(content))
			err = HandleDocument(d)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func HandleDocument(d model.Document) error {
	MyUpload := model.NewDefaultUploader()
	MyUpload.Use(model.DocumentHandler(upload.DBUpdateHandler).ToMiddleWare())

	return MyUpload.Process(d)
}
