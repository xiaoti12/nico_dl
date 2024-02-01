package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"path/filepath"
)

var (
	client     = resty.New()
	cookies    string
	CookieFile string
	MU38File   string
)

func main() {
	loadFileNames()
	loadCookie(CookieFile)
	m3u8URLs := findM3U8(MU38File)
	for i, url := range m3u8URLs {
		//fmt.Println(url)
		m3u8Content := getM3U8(url)
		saveM3U8File(m3u8Content, i)
		url, iv := findKEYAndIV(m3u8Content)
		saveKeyFile(url, iv, i)
		downloadMedia(i, iv)
	}
	mergeMedia()
	fmt.Println("download finished")
}

func loadCookie(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading cookie file: %s\n", err)
	}
	cookies = string(content)
}
func getM3U8(url string) []byte {
	// use resty client to get the content with cookies
	resp, err := client.R().
		SetHeader("Cookie", cookies).
		Get(url)
	if err != nil {
		log.Fatalf("Error getting m3u8: %s\n", err)
	}
	return resp.Body()
}

func loadFileNames() {
	CookieFile = findNameWithExt("txt")
	MU38File = findNameWithExt("m3u8")
}

func findNameWithExt(ext string) string {
	files, err := filepath.Glob("*." + ext)
	if err != nil {
		log.Fatalf("Error searching for %s file: %v\n", ext, err)
	}
	if len(files) == 0 {
		log.Fatalf("No %s file found", ext)
	}
	if len(files) > 1 {
		log.Fatalf("More than one %s file found", ext)
	}
	return files[0]
}
