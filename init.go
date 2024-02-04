package main

import (
	"fmt"
	"log"
	"nico_dl/m3u8"
	"os"
	"path/filepath"
	"strings"
)

func loadData() {
	curPath, _ := os.Getwd()
	log.Println("当前路径:", curPath)
	loadFileNames()
	loadCookie(CookieFile)
	loadCookiesMap()
}

func loadCookie(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("读取cookies文件出错: %s\n", err)
	}
	m3u8.Cookies = string(content)
}

func loadCookiesMap() {
	if m3u8.Cookies == "" {
		fmt.Println("未加载Cookies")
		return
	}
	m3u8.CookiesMap = make(map[string]string)
	kvStrs := strings.Split(m3u8.Cookies, ";")
	for _, s := range kvStrs {
		s = strings.Trim(s, " ")
		kv := strings.Split(s, "=")
		m3u8.CookiesMap[kv[0]] = kv[1]
	}
}

func loadFileNames() {
	CookieFile = findNameWithExt("txt")
	if useM3U8 {
		MU38File = findNameWithExt("m3u8")
	}
}

func findNameWithExt(ext string) string {
	files, err := filepath.Glob("*." + ext)
	if err != nil {
		log.Fatalf("寻找%s文件时出错: %v\n", ext, err)
	}
	if len(files) == 0 {
		log.Fatalf("未找到%s文件", ext)
	}
	if len(files) > 1 {
		log.Fatalf("目录下存在多个%s文件", ext)
	}
	return files[0]
}

func checkRunArgs(args []string) {
	if len(args) == 0 && !useM3U8 {
		log.Fatalf("未提供视频链接")
	}
	if len(args) > 1 {
		log.Fatalf("只能提供一个视频链接")
	}

}
