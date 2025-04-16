package logic

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func IsLocalDevEnv() bool {
	localDev := os.Getenv("LOCAL_DEV_ENV")
	localDevBool, _ := strconv.ParseBool(localDev)
	return localDevBool
}

var DatabaseDirectory string
var AudioFilesDirectory string
var AlbumArtworkDirectory string
var AudioFileTypes []string

func LoadEnv() {

	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "./data"
	}

	audioPath := os.Getenv("AUDIO_PATH")
	if audioPath == "" {
		audioPath = "./audiofiles"
	}

	audioFileTypesEnv := os.Getenv("AUDIO_FILE_TYPES")
	if audioFileTypesEnv == "" {
		AudioFileTypes = []string{
			".aac", ".alac", ".flac", ".m4a", ".mp3", ".ogg",
		}
	} else {
		AudioFileTypes = strings.Split(audioFileTypesEnv, ",")
		// Trim whitespace from each element (optional but recommended)
		for i, ext := range AudioFileTypes {
			AudioFileTypes[i] = strings.TrimSpace(ext)
		}
	}
	log.Printf("Audio file types: %v", AudioFileTypes)

	DatabaseDirectory, _ = filepath.Abs(dataPath)
	AudioFilesDirectory, _ = filepath.Abs(audioPath)
	AlbumArtworkDirectory, _ = filepath.Abs(filepath.Join(dataPath, "album-artwork"))
}

func GenerateSlug() string {
	unixTime := time.Now().Unix()
	unixTimeString := strconv.FormatInt(unixTime, 10)

	nanoTime := time.Now().Nanosecond()
	nanoTimeString := strconv.Itoa(nanoTime)
	return unixTimeString + nanoTimeString
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

func GetDirContents(directoryPath string, fileTypes []string) ([]string, error) {
	if len(fileTypes) > 0 {
		log.Printf("getting contents of %s with file types %v", directoryPath, fileTypes)
	} else {
		log.Printf("getting contents of %s without file type restrictions", directoryPath)
	}
	var foundFiles []string

	absPath, _ := filepath.Abs(directoryPath)

	err := filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error opening directory %s: %s", directoryPath, err)
			return err
		}
		if !info.IsDir() {
			if len(fileTypes) > 0 {
				ext := filepath.Ext(path)
				for _, validExt := range fileTypes {
					if strings.EqualFold(ext, validExt) {
						foundFiles = append(foundFiles, path)
						break
					}
				}
			} else {
				foundFiles = append(foundFiles, path)
			}
		}
		return nil
	})
	log.Printf("Found: %d files in %s", len(foundFiles), directoryPath)
	return foundFiles, err
}

func PrintJsonObject(JsonObject interface{}) {
	jsonData, err := json.MarshalIndent(JsonObject, "", "  ")
	if err != nil {
		log.Printf("Error marshalling JSON: %v\n", err)
	} else {
		log.Printf("JSON Data:\n%s\n", string(jsonData))
	}
}
