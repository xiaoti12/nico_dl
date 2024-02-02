package m3u8

import (
	"fmt"
	"testing"
)

func TestGetAPIData(t *testing.T) {
	Client.SetProxy("http://localhost:7890")
	Cookies = "" // hide for commit
	url := "https://www.nicovideo.jp/watch/sm43273809"
	res := getAPIData(url)
	if res != nil {
		t.Fatalf("Error getting nico website: %s\n", res)
	}

}

func TestGetActiveData(t *testing.T) {
	Client.SetProxy("http://localhost:7890")
	url := "https://www.nicovideo.jp/watch/so43202616"
	data := getActiveData(url)
	if data == "" {
		t.Fatalf("empty data")
	}
	fmt.Println(data)
}
