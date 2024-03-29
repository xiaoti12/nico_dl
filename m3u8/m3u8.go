package m3u8

import (
	"encoding/json"
	"fmt"
	"log"
)

func getM3U8Content(url string) []byte {
	// use resty client to get the content with cookies
	resp, err := Client.R().
		SetHeader("Cookie", Cookies).
		Get(url)
	if err != nil {
		log.Fatalf("远程获取M3U8文件时出错: %s\n", err)
	}
	return resp.Body()
}

func generateM3U8URL(suffix, key, trackID string, audioRate int) string {
	apiURL := fmt.Sprintf("https://nvapi.nicovideo.jp/v1/watch/%s/access-rights/hls?actionTrackId=%s", suffix, trackID)
	postBody := fmt.Sprintf(`{"outputs":[["video-h264-720p","audio-aac-%dkbps"]]}`, audioRate)
	resp, err := Client.R().
		SetHeader("Host", "nvapi.nicovideo.jp").
		SetHeader("Content-Type", "application/json").
		SetHeader("Origin", "https://www.nicovideo.jp").
		SetHeader("Cookie", Cookies).
		SetHeader("X-Access-Right-Key", key).
		SetHeader("X-Frontend-Id", "6").
		SetHeader("X-Request-With", "https://www.nicovideo.jp").
		//SetBody([]byte(`{"outputs":[["video-h264-1080p","audio-aac-128kbps"]]}`)).
		SetBody(postBody).
		Post(apiURL)

	if err != nil {
		log.Println(string(resp.Body()))
		log.Fatalf("远程获取主m3u8文件URL出错: %s\n", err)
	}
	//fmt.Println(string(resp.Body()))
	if resp.Body() == nil {
		log.Fatalf("远程获取主m3u8文件URL的响应为空: %s\n", err)
	}
	return parseM3U8ApiContent(resp.Body())
}
func parseM3U8ApiContent(content []byte) string {
	contentMap := map[string]interface{}{}
	err := json.Unmarshal(content, &contentMap)
	if err != nil {
		log.Fatalf("解析URL内容时出错: %s\n", err)
	}
	dataMap := map[string]interface{}{}
	if _, ok := contentMap["data"]; ok {
		dataMap = contentMap["data"].(map[string]interface{})
		return dataMap["contentUrl"].(string)
	} else {
		log.Fatalf("响应内容不存在主M3U8的URL, 响应内容:%s", string(content))
	}
	return ""
}
