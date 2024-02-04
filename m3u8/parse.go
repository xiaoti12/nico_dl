package m3u8

import (
	"encoding/json"
	"log"
	"strings"
)

func parseData(dataContent string) map[string]interface{} {
	dataMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(dataContent), &dataMap)
	if err != nil {
		log.Fatal("Error parsing nico website data to json", err)
	}
	return dataMap
}

func getAccessKey(dataMap map[string]interface{}) string {
	mediaMap := dataMap["media"].(map[string]interface{})
	domandMap := mediaMap["domand"].(map[string]interface{})
	return domandMap["accessRightKey"].(string)
}

func getTrackID(dataMap map[string]interface{}) string {
	clientMap := dataMap["client"].(map[string]interface{})
	return clientMap["watchTrackId"].(string)
}

func getSupportAudio(dataMap map[string]interface{}) int {
	mediaMap := dataMap["media"].(map[string]interface{})
	domandMap := mediaMap["domand"].(map[string]interface{})
	audiosAttrs := domandMap["audios"].([]interface{})
	supportRate := 0
	for _, attr := range audiosAttrs {
		attrMap := attr.(map[string]interface{})
		id := attrMap["id"].(string)
		if strings.Contains(id, "192") {
			supportRate = max(supportRate, 192)
		}
		if strings.Contains(id, "128") {
			supportRate = max(supportRate, 128)
		}
		if strings.Contains(id, "64") {
			supportRate = max(supportRate, 64)
		}
	}
	return supportRate
}

func getVideoName(dataMap map[string]interface{}) string {
	videoMap := dataMap["video"].(map[string]interface{})
	title := videoMap["title"].(string)
	return title
}
