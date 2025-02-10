package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetConfigDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir = filepath.Join(dir, ".config/gig/")
	return dir, nil
}

func GetCurrentGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	stdout := &strings.Builder{}
	stderr := &strings.Builder{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get current git root. %s %v", stderr, err)
	}
	return strings.Trim(stdout.String(), " \r\n"), nil
}
