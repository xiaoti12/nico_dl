package main

import (
	"fmt"
	"log"
	"nico_dl/m3u8"
	"nico_dl/match"
	"nico_dl/tools"
	"os"
	"path/filepath"
)

var (
	CookieFile string
	MU38File   string
)

func main() {
	loadFileNames()
	loadCookie(CookieFile)
	m3u8URLs := match.FindM3U8URL(MU38File)
	for i, url := range m3u8URLs {
		//fmt.Println(url)
		m3u8Content := getM3U8Content(url)
		m3u8.SaveM3U8File(m3u8Content, i)
		url, iv := match.FindKEYAndIV(m3u8Content)
		m3u8.SaveKeyFile(url, i)
		tools.DownloadMedia(i, iv)
	}
	tools.MergeMedia()
	fmt.Println("download finished")
}

func loadCookie(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading cookie file: %s\n", err)
	}
	m3u8.Cookies = string(content)
}
func getM3U8Content(url string) []byte {
	// use resty client to get the content with cookies
	resp, err := m3u8.Client.R().
		SetHeader("Cookie", m3u8.Cookies).
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
