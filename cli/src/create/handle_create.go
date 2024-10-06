package create

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var DefaultFileName string = "NewBlog"

/*
按照模板创建新的文件

不再设置中间件处理新建的文件，而是采用直接修改模板

dir: 指定的文件路径的文件夹
*/
func HandleCreateFile(dir string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("捕获到 panic: %v", r)
		}
	}()

	// 获取新建文件obj
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	newFileObj := GetNewFileObj(dir)

	// 获取模板obj
	exeName, err := os.Executable()
	if err != nil {
		return err
	}
	templateFileName := fmt.Sprintf("%s/default.template", filepath.Dir(exeName))
	templateFileObj := GetTemplateFileObj(templateFileName)

	_, err = io.Copy(newFileObj, templateFileObj)
	if err != nil {
		return err
	}

	return nil
}

func GetNewFileObj(dir string) *os.File {
	var fileObj *(os.File)
	var err error

	NewFileName := fmt.Sprintf("%s/%s.md", dir, DefaultFileName)
	if _, err = os.Stat(NewFileName); os.IsNotExist(err) {
		fileObj, err = os.Create(NewFileName)

		if err != nil {
			panic(err)
		}
	} else {
		count := 1
		for {
			NewFileName = fmt.Sprintf("%s/%s_%d.md", dir, DefaultFileName, count)
			_, err = os.Stat(NewFileName)
			if os.IsNotExist(err) {
				// 不存在文件,可以创建
				fileObj, err = os.Create(NewFileName)
				if err != nil {
					panic(err)
				}
				break
			} else if err != nil {
				// 其他Err
				panic(err)
			}

			count += 1
			if count > 32 {
				panic("try too many times while creating new file")
			}
		}
	}
	return fileObj
}

func GetTemplateFileObj(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}
