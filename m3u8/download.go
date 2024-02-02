package m3u8

import (
	"fmt"
	"nico_dl/match"
	"nico_dl/tools"
)

func DLMediasWithM3U8File(topFile string) {
	// video and audio m3u8 urls
	urls := match.FindURLsWithFile(topFile)
	for i, url := range urls {
		m3u8Content := getM3U8Content(url)
		saveM3U8File(m3u8Content, i)
		keyURL, iv := match.FindKEYAndIV(m3u8Content)
		saveKeyFile(keyURL, i)
		tools.DownloadMedia(i, iv)
	}
	return
}

func DLMediaWithCode(suffix string) {
	nicoURL := fmt.Sprintf("https://www.nicovideo.jp/watch/%s", suffix)

	apiData := getAPIData(nicoURL)
	apiDataMap := parseData(apiData)

	accessKey := getAccessKey(apiDataMap)
	trackID := getTrackID(apiDataMap)
	audioRate := getSupportAudio(apiDataMap)
	fmt.Println("accessKey:", accessKey)
	fmt.Println("trackID:", trackID)

	m3u8URL := generateM3U8URL(suffix, accessKey, trackID, audioRate)
	m3u8Content := getM3U8Content(m3u8URL)
	// video and audio m3u8 urls
	urls := match.FindURLs(m3u8Content)
	for i, url := range urls {
		m3u8Content := getM3U8Content(url)
		saveM3U8File(m3u8Content, i)
		keyURL, iv := match.FindKEYAndIV(m3u8Content)
		saveKeyFile(keyURL, i)
		tools.DownloadMedia(i, iv)
	}
}
