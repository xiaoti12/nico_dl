package m3u8

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"nico_dl/match"
	"nico_dl/tools"
	"os"
)

var (
	Client     = resty.New()
	Cookies    string
	CookiesMap map[string]string
)

func SaveKeyFile(keyURL string, suffix int) {
	resp, err := Client.R().
		SetHeader("Cookie", Cookies).
		Get(keyURL)
	if err != nil {
		fmt.Printf("Error getting content: %s\n", err)
		return
	}
	fileName := fmt.Sprintf("key_%d.key", suffix)
	err = os.WriteFile(fileName, resp.Body(), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return
	}
	fmt.Printf("Key file saved to: %s\n", fileName)
}

func saveM3U8File(content []byte, suffix int) {
	fileName := fmt.Sprintf("m3u8_%d.m3u8", suffix)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return
	}
	fmt.Printf("M3U8 file saved to: %s\n", fileName)
}

func DLMediaFiles(topFile string) {
	urls := match.FindM3U8URL(topFile)
	for i, url := range urls {
		m3u8Content := getM3U8Content(url)
		saveM3U8File(m3u8Content, i)
		keyURL, iv := match.FindKEYAndIV(m3u8Content)
		SaveKeyFile(keyURL, i)
		tools.DownloadMedia(i, iv)
	}
	return
}
