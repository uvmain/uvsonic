package logic

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func IsLocalDevEnv() bool {
	localDev := os.Getenv("LOCAL_DEV_ENV")
	localDevBool, _ := strconv.ParseBool(localDev)
	return localDevBool
}

func GenerateSlug() string {
	unixTime := time.Now().Unix()
	unixTimeString := strconv.FormatInt(unixTime, 10)

	nanoTime := time.Now().Nanosecond()
	nanoTimeString := strconv.Itoa(nanoTime)
	return unixTimeString + nanoTimeString
}

var DatabaseDirectory string
var AudioFilesDirectory string
var AlbumArtworkDirectory string

func LoadEnv() {

	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "./data"
	}

	audioPath := os.Getenv("AUDIO_PATH")
	if audioPath == "" {
		audioPath = "./audiofiles"
	}

	DatabaseDirectory, _ = filepath.Abs(dataPath)
	AudioFilesDirectory, _ = filepath.Abs(audioPath)
	AlbumArtworkDirectory, _ = filepath.Abs(filepath.Join(dataPath, "album-artwork"))
}

func CreateDir(directoryPath string) {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		log.Printf("Creating directory: %s", directoryPath)
		err := os.MkdirAll(directoryPath, 0755)
		if err != nil {
			log.Printf("Error creating directory%s: %s", directoryPath, err)
		} else {
			log.Printf("Directory created: %s", directoryPath)
		}
	} else {
		log.Printf("Directory already exists: %s", directoryPath)
	}
}

func GetDirContents(directoryPath string) ([]string, error) {
	var foundFiles []string

	absPath, _ := filepath.Abs(directoryPath)

	err := filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error opening directory %s: %s", directoryPath, err)
			return err
		}
		if !info.IsDir() {
			foundFiles = append(foundFiles, path)
		}
		return nil
	})
	log.Printf("Found: %d images in %s", len(foundFiles), directoryPath)
	return foundFiles, err
}
