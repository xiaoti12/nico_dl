package m3u8

import (
	"encoding/json"
	"log"
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
