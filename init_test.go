package main

import (
	"nico_dl/m3u8"
	"testing"
)

func TestLoadFileNames(t *testing.T) {
	loadFileNames()
	if CookieFile == "" {
		t.Fatalf("No cookie file found")
	}
	t.Log(MU38File)
	t.Log(CookieFile)
}
func TestLoadCookie(t *testing.T) {
	TestLoadFileNames(t)
	loadCookie(CookieFile)
	if m3u8.Cookies == "" {
		t.Fatalf("No cookie loaded")
	}
	t.Log(m3u8.Cookies)
}
func TestLoadCookiesMap(t *testing.T) {
	TestLoadCookie(t)
	loadCookiesMap()
	if m3u8.CookiesMap == nil || len(m3u8.CookiesMap) == 0 {
		t.Fatalf("No cookie loaded")
	}
	t.Log(m3u8.CookiesMap)
}
