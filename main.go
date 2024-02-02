package main

import (
	"fmt"
	"nico_dl/m3u8"
	"nico_dl/tools"
)

var (
	CookieFile string
	MU38File   string
)

func main() {
	loadData()
	m3u8.DLMediaFiles(MU38File)
	tools.MergeMedia()
	fmt.Println("download finished")
}
