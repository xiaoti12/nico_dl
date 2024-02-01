package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func downloadMedia(fileSuffix int, iv string) {
	curPath, _ := os.Getwd()
	m3u8Name := fmt.Sprintf("m3u8_%d.m3u8", fileSuffix)
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
	runShellCommand(cmdStr)
	if err := os.Remove(m3u8Name); err != nil {
		fmt.Println("remove m3u8 file failed:", err)
	}
	if err := os.Remove(keyFileName); err != nil {
		fmt.Println("remove key file failed:", err)
	}
}

func mergeMedia() {
	cmdArgs := []string{
		"ffmpeg.exe",
		"-i", "m3u8_0.m4a",
		"-i", "m3u8_1.mp4",
		"-c", "copy",
		"nicovideo.mp4",
	}
	// ffmpeg error if the output file exists
	if _, err := os.Stat("nicovideo.mp4"); err == nil {
		os.Remove("nicovideo.mp4")
	}
	cmdStr := strings.Join(cmdArgs, " ")
	runShellCommand(cmdStr)
	if err := os.Remove("m3u8_0.m4a"); err != nil {
		fmt.Println("remove audio file failed:", err)
	}
	if err := os.Remove("m3u8_1.mp4"); err != nil {
		fmt.Println("remove video file failed:", err)
	}
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
		log.Fatalf("run program failed: %v", err)
	}
}
