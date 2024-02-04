package main

import (
	"log"
	"nico_dl/m3u8"
	"nico_dl/tools"
)

var (
	CookieFile string
	MU38File   string
)

func main() {
	loadData()
	//m3u8.DLMediasWithM3U8File(MU38File)
	m3u8.DLMediaWithCode("sm43273809")
	tools.MergeMedia()
	log.Println("下载完成")
}
