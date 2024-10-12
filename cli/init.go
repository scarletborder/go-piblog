package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var exeDir string = ""

const (
	templateFileName string = "default.template"
)

func init() {
	// 获取可执行文件的路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取可执行文件所在的目录
	exeDir = filepath.Dir(exePath)
	GuaranteeTemplate()

}

const templateFile = `# Example Title
[tag ex1, tag ex2]
The text following tags will be recognized as brief text.  

Brief text ends when meet '---'.  
---

Some texts which only display in main text  

## Example subtitle 1

Some texts which only display in main text

## Example subtitle 2

### Some text

Some texts which only display in main text`

func GuaranteeTemplate() {
	templatePath := fmt.Sprintf("%s/%s", exeDir, templateFileName)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		fmt.Println("第一次创建模板文件")
		obj, err := os.Create(templatePath)
		if err != nil {
			panic(err)
		}
		_, err = obj.WriteString(templateFile)
		if err != nil {
			panic(err)
		}
	}
}
