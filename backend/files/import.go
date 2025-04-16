package files

import (
	"github.com/uvmain/uvsonic/ff"
	"github.com/uvmain/uvsonic/logic"
)

func Init() {
	importFiles()
}

func importFiles() {
	filePaths, _ := logic.GetDirContents(logic.AudioFilesDirectory, logic.AudioFileTypes)
	for _, filePath := range filePaths {
		go ff.GetTags(filePath)
	}
}
