package upload

import "os"

// forceNew: 是否每次创建新的博文
func HandleUploadBlog(filePath string, forceNew bool) (err error) {
	// TODO: Get file obj in specified path
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	// TODO: Read file and convert to string
	// To struct document
	original_text := string(fileBytes)
	document := ParseDocument(original_text)

	// TODO: Process all middleware
	if forceNew {

	} else {
		MyUpload := GlobalUpload.Copy()
		MyUpload.Use(DocumentHandler(DBUpdateHandler).ToMiddleWare())
		err = MyUpload.Process(document)
	}

	return
}
