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

func download(fileSuffix int, iv string) {
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
	//fmt.Println(cmdStr)
	cmd := exec.Command("cmd.exe", "/c", cmdStr)
	cmd.Dir = curPath
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("set stdout failed: %v", err)
	}
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	err = cmd.Run()
	if err != nil {
		log.Fatalf("run download program failed: %v", err)
	}
}
