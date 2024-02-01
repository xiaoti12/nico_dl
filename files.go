package main

import (
	"fmt"
	"os"
)

func saveKeyFile(keyURL, iv string, suffix int) {
	resp, err := client.R().
		SetHeader("Cookie", cookies).
		Get(keyURL)
	if err != nil {
		fmt.Printf("Error getting content: %s\n", err)
		return
	}
	fileName := fmt.Sprintf("key_%d.key", suffix)
	err = os.WriteFile(fileName, resp.Body(), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return
	}
	fmt.Printf("Key file saved to: %s\n", fileName)
}

func saveM3U8File(content []byte, suffix int) {
	fileName := fmt.Sprintf("m3u8_%d.m3u8", suffix)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return
	}
	fmt.Printf("M3U8 file saved to: %s\n", fileName)
}
