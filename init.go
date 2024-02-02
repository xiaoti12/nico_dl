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
	loadFileNames()
	loadCookie(CookieFile)
	loadCookiesMap()
}

func loadCookie(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading cookie file: %s\n", err)
	}
	m3u8.Cookies = string(content)
}

func loadCookiesMap() {
	if m3u8.Cookies == "" {
		fmt.Println("No cookie found")
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
