package videoutils

import (
	"log"
	"regexp"
	"strconv"
)

type PlaylistItem struct {
	Width     int
	Height    int
	Bandwidth int
	File      string
}

func NewItem(itemString string) *PlaylistItem {
	item := new(PlaylistItem)

	regExpItem, _ := regexp.Compile(`([A-Z]+)=([^,\n]+)`)
	regExpFile, _ := regexp.Compile(`\n([^\.]+\.m3u8)`)
	item.File = regExpFile.FindAllStringSubmatch(itemString, 1)[0][1]

	subMatch := regExpItem.FindAllStringSubmatch(itemString, -1)
	for _, v := range subMatch {
		switch v[1] {
		case "BANDWIDTH":
			item.Bandwidth, _ = strconv.Atoi(v[2])
			break
		case "RESOLUTION":
			resRegExp, _ := regexp.Compile(`x`)
			resolution := resRegExp.Split(v[2], 2)
			item.Width, _ = strconv.Atoi(resolution[0])
			item.Height, _ = strconv.Atoi(resolution[1])
			break
		}
	}
	return item
}

func xx(str string) {
	log.Println(str)
}
