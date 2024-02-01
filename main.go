package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

var (
	client  = resty.New()
	cookies string
)
var (
	CookieFile = "nico_cookie.txt"
	MU38File   = "entering.m3u8"
)

func main() {
	loadCookie(CookieFile)
	m3u8URLs := findM3U8(MU38File)
	for i, url := range m3u8URLs {
		//fmt.Println(url)
		m3u8Content := getM3U8(url)
		saveM3U8File(m3u8Content, i)
		url, iv := findKEYAndIV(m3u8Content)
		saveKeyFile(url, iv, i)
		download(i, iv)
	}
	//download(0, "0x00")
}

func loadCookie(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading cookie file: %s\n", err)
		return
	}
	cookies = string(content)
}
func getM3U8(url string) []byte {
	// use resty client to get the content with cookies
	resp, err := client.R().
		SetHeader("Cookie", cookies).
		Get(url)
	if err != nil {
		fmt.Printf("Error getting m3u8: %s\n", err)
		return nil
	}
	fmt.Println("download m3u8 file")
	return resp.Body()
}
