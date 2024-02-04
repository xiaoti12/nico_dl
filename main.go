package main

import (
	"flag"
	"log"
	"nico_dl/m3u8"
	"nico_dl/tools"
	"path/filepath"
)

var (
	CookieFile string
	MU38File   string
)
var (
	useM3U8  bool
	videoURL string
)

func main() {
	flag.BoolVar(&useM3U8, "m3u8", false, "使用m3u8文件")
	flag.Parse()

	args := flag.Args()
	checkRunArgs(args)
	var suffix string
	if !useM3U8 {
		videoURL = args[0]
		suffix = filepath.Base(videoURL)
	}

	loadData()
	if useM3U8 {
		m3u8.DLMediasWithM3U8File(MU38File)
	} else {
		m3u8.DLMediaWithCode(suffix)
	}

	fileName := tools.MergeMedia(suffix, m3u8.VideoName)
	log.Printf("视频下载完成：%s\n", fileName)
}
