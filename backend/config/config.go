package config

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var FFMPEG_PATH string
var FFPROBE_PATH string

func LoadConfig() {

	ffmpegPath := os.Getenv("FFMPEG_PATH")
	if ffmpegPath == "" {
		ffmpegPath = "lib/ffmpeg"
	}
	FFMPEG_PATH, _ = filepath.Abs(ffmpegPath)

	log.Printf("FFMPEG_PATH: %s", FFMPEG_PATH)
	cmd := exec.Command(FFMPEG_PATH, "-version")
	err := cmd.Run()
	if err != nil {
		log.Printf("ffmpeg not found at %s", FFMPEG_PATH)
	} else {
		log.Printf("ffmpeg found at %s", FFMPEG_PATH)
	}

	ffprobePath := os.Getenv("FFPROBE_PATH")
	if ffprobePath == "" {
		ffprobePath = "lib/ffprobe"
	}
	FFPROBE_PATH, _ = filepath.Abs(ffprobePath)

	log.Printf("FFPROBE_PATH: %s", FFPROBE_PATH)
	cmd = exec.Command(FFPROBE_PATH, "-version")
	err = cmd.Run()
	if err != nil {
		log.Printf("ffprobe not found at %s", FFPROBE_PATH)
	} else {
		log.Printf("ffprobe found at %s", FFPROBE_PATH)
	}
}
