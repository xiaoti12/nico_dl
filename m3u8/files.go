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
	VideoName  string
)

func saveKeyFile(keyURL string, suffix int) {
	resp, err := Client.R().
		SetHeader("Cookie", Cookies).
		Get(keyURL)
	if err != nil {
		log.Fatalf("远程获取key文件时出错: %s\n", err)
	}
	fileName := fmt.Sprintf("key_%d.key", suffix)
	err = os.WriteFile(fileName, resp.Body(), 0644)
	if err != nil {
		log.Fatalf("保存key文件%s时出错: %s\n", fileName, err)
	}
	log.Printf("保存Key文件为: %s\n", fileName)
}

func saveM3U8File(content []byte, suffix int) {
	fileName := fmt.Sprintf("m3u8_%d.m3u8", suffix)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		log.Fatalf("保存M3U8文件%s时出错: %s\n", fileName, err)
	}
	log.Printf("保存M3U8文件为: %s\n", fileName)
}
