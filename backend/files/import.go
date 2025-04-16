package files

import (
	"log"

	"github.com/uvmain/uvsonic/db"
	"github.com/uvmain/uvsonic/ff"
	"github.com/uvmain/uvsonic/logic"
	"github.com/uvmain/uvsonic/types"
)

var metadataChannel = make(chan types.TrackMetadata, 100)

func Init() {
	go startMetadataWorker()
	importFiles()
}

func importFiles() {
	filePaths, _ := logic.GetDirContents(logic.AudioFilesDirectory, logic.AudioFileTypes)
	for _, filePath := range filePaths {
		go processFile(filePath)
	}
}

func processFile(filepath string) {
	metadata, err := ff.GetTags(filepath)
	if err != nil {
		log.Printf("Failed to get tags for file %s: %s", filepath, err)
		return
	}

	metadataChannel <- metadata
}

func startMetadataWorker() {
	for metadata := range metadataChannel {
		err := db.InsertTrackMetadata(metadata)
		if err != nil {
			log.Printf("Failed to insert metadata for file %s: %s", metadata.Filename, err)
		}
	}
}
