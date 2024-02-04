package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DownloadMedia(fileSuffix int, iv string) {
	//如果路径存在非法字符时会报错
	curPath, _ := os.Getwd()
	m3u8Name := filepath.Join(curPath, fmt.Sprintf("m3u8_%d.m3u8", fileSuffix))
	saveName := fmt.Sprintf("m3u8_%d", fileSuffix)
	keyFileName := filepath.Join(curPath, fmt.Sprintf("key_%d.key", fileSuffix))

	cmdArgs := []string{
		"chcp 65001 &",
		"N_m3u8DL-CLI_v3.0.2.exe", m3u8Name,
		"--saveName", saveName,
		"--workDir", curPath,
		"--useKeyFile", keyFileName,
		"--useKeyIV", iv,
		"--enableDelAfterDone",
	}
	cmdStr := strings.Join(cmdArgs, " ")
	fmt.Println("==========开始使用n_m3u8DL-CLI下载媒体==========")
	runShellCommand(cmdStr)
	fmt.Println("==========n_m3u8DL-CLI下载媒体完成==========")
	if err := os.Remove(m3u8Name); err != nil {
		fmt.Printf("移除%s出错:%s\n", m3u8Name, err)
	}
	if err := os.Remove(keyFileName); err != nil {
		fmt.Printf("移除%s出错:%s\n", keyFileName, err)
	}
}

func MergeMedia(suffix, name string) string {
	var fullFileName string
	if name == "" {
		fullFileName = fmt.Sprintf("nicovideo_%s.mp4", suffix)
	} else {
		fullFileName = fmt.Sprintf(`%s.mp4`, name)
	}
	//fmt.Println(fullFileName)
	cmdArgs := []string{
		"ffmpeg.exe",
		"-i", "m3u8_0.m4a",
		"-i", "m3u8_1.mp4",
		"-c", "copy",
		fullFileName,
	}
	if _, err := os.Stat("m3u8_0.m4a"); err != nil {
		log.Fatal("未找到音频文件m3u8_0.m4a")
	}
	if _, err := os.Stat("m3u8_1.mp4"); err != nil {
		log.Fatal("未找到视频文件m3u8_1.mp4")
	}
	// ffmpeg error if the output file exists
	if _, err := os.Stat(fullFileName); err == nil {
		os.Remove(fullFileName)
	}
	cmdStr := strings.Join(cmdArgs, " ")
	runShellCommand(cmdStr)
	if err := os.Remove("m3u8_0.m4a"); err != nil {
		fmt.Printf("移除音频文件出错:%s\n", err)
	}
	if err := os.Remove("m3u8_1.mp4"); err != nil {
		fmt.Printf("移除视频文件出错:%s\n", err)
	}
	return fullFileName
}

func runShellCommand(cmdStr string) {
	cmd := exec.Command("cmd.exe", "/c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("set stdout failed: %v", err)
	}
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		stdout.Close()
	}()
	err = cmd.Run()
	if err != nil {
		log.Fatalf("运行外部程序时出错: %v", err)
	}
}
