package m3u8

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

var (
	Client     = resty.New()
	Cookies    string
	CookiesMap map[string]string
)

func saveKeyFile(keyURL string, suffix int) {
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
	log.Printf("保存Key文件为: %s\n", fileName)
}

func saveM3U8File(content []byte, suffix int) {
	fileName := fmt.Sprintf("m3u8_%d.m3u8", suffix)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %s\n", err)
	}
	log.Printf("保存M3U8文件为: %s\n", fileName)
}
