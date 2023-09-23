package config

import (
	"fmt"
	"os/exec"
	"regexp"
)

const configDir string = "$HOME/.gid"

func GetConfigDir() string {
	out, err := exec.Command("bash", "-c", fmt.Sprintf("echo %v", configDir)).Output()
	if err != nil {
		return ""
	}

	dirOut := regexp.MustCompile(`[^a-zA-Z0-9\./\\]+`).ReplaceAllString(string(out), "")
	return dirOut
}

var ConfigDir string = GetConfigDir()
var ConfigPath string = fmt.Sprintf("%v/.config", ConfigDir)
