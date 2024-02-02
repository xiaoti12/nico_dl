package m3u8

import (
	"testing"
)

func TestGetAPIData(t *testing.T) {
	Client.SetProxy("http://localhost:7890")
	Cookies = "" // hide for commit
	url := "https://www.nicovideo.jp/watch/sm43273809"
	res := getAPIData(url)
	if res != "" {
		t.Fatalf("Error getting nico website: %s\n", res)
	}
	t.Log(res)
}
