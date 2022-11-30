package videoutils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pjmd89/goutils/systemutils"
)

const (
	rate1080       = "4500k"
	minRate1080    = "4500k"
	maxRate1080    = "9000k"
	bufferSize1080 = "9000k"
	scale1080      = "-1:1080"
	path1080       = "1080"

	rate720       = "2500k"
	minRate720    = "1500k"
	maxRate720    = "4000k"
	bufferSize720 = "5000k"
	scale720      = "-1:720"
	path720       = "720"

	rate480       = "1000k"
	minRate480    = "500k"
	maxRate480    = "2000k"
	bufferSize480 = "2000k"
	scale480      = "-1:480"
	path480       = "480"

	rate360       = "750k"
	minRate360    = "400k"
	maxRate360    = "1000k"
	bufferSize360 = "1500k"
	scale360      = "-1:360"
	path360       = "360"

	rate240       = "500k"
	minRate240    = "300k"
	maxRate240    = "750k"
	bufferSize240 = "1000k"
	scale240      = "-1:240"
	path240       = "240"

	rateOption       = "-b:v"
	minRateOption    = "-minrate"
	maxRateOption    = "-maxrate"
	bufferSizeOption = "-bufsize"
	scaleOption      = "-vf"
)

type FFMpeg struct {
	path       string
	height     int
	width      int
	fileName   string
	playlist   string
	master     string
	Status     chan status
	done       bool
	resolution map[string]resolution
	filePath   string
	newPath    string
	splitPath  string
}

type status struct {
	Speed   string
	Frames  int
	Time    string
	Bitrate int
}
type resolution struct {
	rate       string
	minRate    string
	maxRate    string
	bufferSize string
	scale      string
	path       string
}

func Create(path string, height int, width int) *FFMpeg {
	resolutions := map[string]resolution{
		"1080": {rate1080, minRate1080, maxRate1080, bufferSize1080, scale1080, path1080},
		"720":  {rate720, minRate720, maxRate720, bufferSize720, scale720, path720},
		"480":  {rate480, minRate480, maxRate480, bufferSize480, scale480, path480},
		"360":  {rate360, minRate360, maxRate360, bufferSize360, scale360, path360},
		"240":  {rate240, minRate240, maxRate240, bufferSize240, scale240, path240},
	}
	return &FFMpeg{
		path,
		height,
		width,
		"",
		"",
		"",
		make(chan status),
		false,
		resolutions,
		"",
		"",
		"",
	}
}
func (ff *FFMpeg) directoriesPrepair() error {
	filePath, _ := filepath.Abs(ff.path)
	if !systemutils.FileExists(filePath) {
		return fmt.Errorf("[ffmpeg] File " + filePath + " do not exists")
	}
	dirPath := filepath.Dir(filePath)
	baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	newPath := dirPath + "/" + baseName
	splitPath := newPath + "/" + strconv.Itoa(ff.height)

	if !systemutils.FileExists(splitPath) {
		if err := os.MkdirAll(splitPath, os.ModePerm); err != nil {
			return fmt.Errorf("[ffmpeg] folder " + splitPath + " has not created(splitPath)")
		}
	}
	ff.filePath = filePath
	ff.newPath = newPath
	ff.splitPath = splitPath
	return nil
}
func (ffmpeg *FFMpeg) Transcoder() bool {
	if err := ffmpeg.directoriesPrepair(); err != nil {
		log.Fatal(err.Error())
		return false
	}
	arguments := ffmpeg.setArguments(false)
	cmd := exec.Command("ffmpeg", arguments...)
	cmdReader, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(cmdReader)
	scanner.Split(splitScan)

	go ffmpeg.transcoderScan(scanner)
	cmd.Start()
	return true
}
func (ff *FFMpeg) HLS() bool {
	if err := ff.directoriesPrepair(); err != nil {
		log.Fatal(err.Error())
		return false
	}
	arguments := ff.setArguments(true)
	cmd := exec.Command("ffmpeg", arguments...)
	cmd.Output()
	return true
}
func (ff *FFMpeg) IsDone() bool {

	return ff.done
}
func (ff *FFMpeg) GetThumb() {
	arguments := ff.setThumbArguments()
	cmd := exec.Command("ffmpeg", arguments...)
	cmd.Output()
}
func (ff *FFMpeg) setThumbArguments() (r []string) {
	arguments := []string{
		"-i",
		ff.filePath,
		"-ss",
		"00:00:15",
		"-frames:v",
		"1",
		ff.splitPath + "/thumb.jpg",
	}
	r = arguments
	return
}
func (ff *FFMpeg) setArguments(hls bool) []string {
	var arguments []string
	/*
		arguments = []string{
			"-i",
			ff.filePath,
			"-y",
			"-vcodec",
			"libx265",
			"-tag:v",
			"hvc1",
			"-crf",
			"28",
			"-vf",
			"scale=" + strconv.Itoa(ff.width) + ":" + strconv.Itoa(ff.height), // + ":force_original_aspect_ratio=decrease",
			"-acodec",
			"copy",
			"-f",
			"mp4",
			"-hide_banner",
			"-preset",
			"fast",
		}
	*/
	//"-y -vcodec libx264 -tag:v avc1 -strict -2 -crf 28 -acodec aac -f mp4 -hide_banner -preset ultrafast"
	arguments = []string{
		"-i",
		ff.filePath,
		"-y",
		"-vcodec",
		"libx264",
		"-tag:v",
		"avc1",
		"-strict",
		"-2",
		"-crf",
		"28",
		"-vf",
		"scale=" + strconv.Itoa(ff.width) + ":" + strconv.Itoa(ff.height), // + ":force_original_aspect_ratio=decrease",
		"-acodec",
		"aac",
		"-f",
		"mp4",
		"-hide_banner",
		"-preset",
		"ultrafast",
	}
	ff.fileName = ff.splitPath + "/output.mp4"
	arguments = append(arguments, ff.splitPath+"/output.mp4")
	if hls {
		arguments = []string{
			"-i",
			ff.splitPath + "/output.mp4",
			"-y",
			"-c",
			"copy",
			"-hls_time",
			"10",
			"-hls_playlist_type",
			"vod",
			"-hls_flags",
			"delete_segments+round_durations+split_by_time",
			"-master_pl_name",
			"master.m3u8",
		}
		arguments = append(arguments, ff.splitPath+"/output.m3u8")

		ff.playlist = ff.splitPath + "/output.m3u8"
		ff.master = ff.splitPath + "/master.m3u8"
	}
	//fmt.Println(strings.Join(arguments[:]," "))
	return arguments
}
func (ff *FFMpeg) DeleteFile() bool {
	err := os.Remove(ff.fileName)
	response := true
	if err != nil {
		response = false
	}
	return response
}
func (ff *FFMpeg) GetFileName() string {
	return ff.fileName
}
func (ff *FFMpeg) GetPlaylist() string {
	return ff.playlist
}
func (ff *FFMpeg) GetMasterPlaylist() string {
	return ff.master
}
func (ff *FFMpeg) transcoderScan(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		//log.Println(line)

		if strings.Contains(line, "frame=") && strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
			var re = regexp.MustCompile(`=\s+`)
			st := re.ReplaceAllString(line, `=`)

			f := strings.Fields(st)
			var framesProcessed string
			var currentTime string
			var currentBitrate string
			var currentSpeed string

			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")

				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]

					if fieldname == "frame" {
						framesProcessed = fieldvalue
					}

					if fieldname == "time" {
						currentTime = fieldvalue
					}

					if fieldname == "bitrate" {
						currentBitrate = fieldvalue
					}
					if fieldname == "speed" {
						currentSpeed = fieldvalue
					}
				}
			}
			frames, _ := strconv.Atoi(framesProcessed)
			bitrate, _ := strconv.Atoi(currentBitrate)
			if frames <= 0 {
				frames = 1
			}
			stat := status{currentSpeed, frames, currentTime, bitrate}
			select {
			case ff.Status <- stat:
			default:
			}
		}
	}
	ff.done = true
	close(ff.Status)
}
func splitScan(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	if i := bytes.IndexByte(data, '\r'); i >= 0 {
		// We have a cr terminated line
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
