package videoutils

import (
	"encoding/xml"
	"os/exec"

	"github.com/pjmd89/goutils/systemutils"
)

type MediaInfoXML struct {
	MediaInfo xml.Name `xml:"MediaInfo"`
	Media     struct {
		Ref   string `xml:"ref,attr"`
		Track []struct {
			Type                 string `xml:"type,attr"`
			FrameCount           int
			Width                int
			Height               int
			Duration             float64
			Encoded_Library_Name string
			Format               string
		} `xml:"track"`
	} `xml:"media"`
}
type MediaInfo struct {
	Track map[string]track
}
type track struct {
	FrameCount         int
	Width              int
	Height             int
	Duration           float64
	EncodedLibraryName string
	Format             string
}

func GetMediaInfo(filepath *systemutils.FileInfo) MediaInfo {
	arguments := []string{
		"--Output=XML",
		filepath.Abs,
	}
	cmd := exec.Command("mediainfo", arguments...)
	out, _ := cmd.Output()
	mediaInfoXML := &MediaInfoXML{}
	mediaInfo := &MediaInfo{}
	xml.Unmarshal(out, mediaInfoXML)
	mediaInfo.Track = make(map[string]track)
	for _, val := range mediaInfoXML.Media.Track {
		trackStruct := track{val.FrameCount, val.Width, val.Height, val.Duration, val.Encoded_Library_Name, val.Format}
		mediaInfo.Track[val.Type] = trackStruct

	}
	return *mediaInfo
}
