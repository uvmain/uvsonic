package ff

import (
	"log"
	"os/exec"

	"github.com/uvmain/uvsonic/config"
)

func GetTags(audfilePath string) (interface{}, error) {
	cmd := exec.Command(config.FFPROBE_PATH, "-show_format", "-print_format", "json", audfilePath)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error running ffprobe: %v", err)
		return nil, err
	}
	log.Printf("Tags: %s", string(output))
	return output, nil
}
