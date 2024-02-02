package match

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func FindM3U8URL(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return nil
	}
	pattern := `http[s]?://[^\s"]+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(string(content), -1)
	return matches
}

func findMatch(content []byte, prefix string, getIV bool) (string, string) {
	lines := strings.Split(string(content), "\n")
	var pattern string
	var url, iv string
	if getIV {
		pattern = `URI="([^"]+)",IV=([^,]+)`
	} else {
		pattern = `URI="([^"]+)"`
	}
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			re := regexp.MustCompile(pattern)
			match := re.FindStringSubmatch(line)
			url = match[1]
			if getIV {
				iv = match[2]
			}
		}
	}
	return url, iv
}

func FindKEYAndIV(content []byte) (string, string) {
	lines := strings.Split(string(content), "\n")
	var key, iv string

	for _, line := range lines {
		// 查找 #EXT-X-KEY 字段
		if strings.HasPrefix(line, "#EXT-X-KEY:") {
			re := regexp.MustCompile(`URI="([^"]+)",IV=([^,]+)`)
			match := re.FindStringSubmatch(line)
			if len(match) == 3 {
				key = match[1]
				iv = match[2]
				break
			}
		}
	}
	iv = iv[2:]
	return key, iv
}
