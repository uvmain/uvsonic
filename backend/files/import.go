package files

import (
	"log"

	"github.com/uvmain/uvsonic/db"
	"github.com/uvmain/uvsonic/ff"
	"github.com/uvmain/uvsonic/logic"
)

func Init() {
	importFiles()
}

func importFiles() {
	filePaths, _ := logic.GetDirContents(logic.AudioFilesDirectory, logic.AudioFileTypes)
	for _, filePath := range filePaths {
		go importFile(filePath)
	}
}

func importFile(filepath string) {
	metadata, err := ff.GetTags(filepath)
	if err != nil {
		log.Printf("Failed to get tags for file %s: %s", filepath, err)
		return
	}

	err = db.InsertTrackMetadata(metadata)
	if err != nil {
		log.Printf("Failed to get tags for file: %s", err)
		return
	}
}
