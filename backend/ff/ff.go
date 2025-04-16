package ff

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/uvmain/uvsonic/config"
	"github.com/uvmain/uvsonic/types"
)

func GetTags(audfilePath string) (types.TrackMetadata, error) {
	cmd := exec.Command(config.FFPROBE_PATH, "-show_format", "-print_format", "json", audfilePath)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error running ffprobe: %v", err)
		return types.TrackMetadata{}, err
	}

	var ffprobeOutput struct {
		Format struct {
			Filename   string            `json:"filename"`
			FormatName string            `json:"format_name"`
			Tags       map[string]string `json:"tags"`
			Duration   string            `json:"duration"`
			Size       string            `json:"size"`
			Bitrate    string            `json:"bit_rate"`
		} `json:"format"`
	}

	err = json.Unmarshal(output, &ffprobeOutput)
	if err != nil {
		log.Printf("Error parsing ffprobe output: %v", err)
		return types.TrackMetadata{}, err
	}

	metadata := types.TrackMetadata{
		Filename:            ffprobeOutput.Format.Filename,
		Format:              ffprobeOutput.Format.FormatName,
		Duration:            ffprobeOutput.Format.Duration,
		Size:                ffprobeOutput.Format.Size,
		Bitrate:             ffprobeOutput.Format.Bitrate,
		Title:               ffprobeOutput.Format.Tags["TITLE"],
		Artist:              ffprobeOutput.Format.Tags["ARTIST"],
		Album:               ffprobeOutput.Format.Tags["ALBUM"],
		AlbumArtist:         ffprobeOutput.Format.Tags["album_artist"],
		Genre:               ffprobeOutput.Format.Tags["GENRE"],
		TrackNumber:         ffprobeOutput.Format.Tags["track"],
		DiscNumber:          ffprobeOutput.Format.Tags["disc"],
		ReleaseDate:         ffprobeOutput.Format.Tags["DATE"],
		MusicBrainzArtistID: ffprobeOutput.Format.Tags["MUSICBRAINZ_ARTISTID"],
		MusicBrainzAlbumID:  ffprobeOutput.Format.Tags["MUSICBRAINZ_ALBUMID"],
		Label:               ffprobeOutput.Format.Tags["LABEL"],
	}

	log.Printf("Metadata read: %+v", metadata)

	return metadata, nil
}
