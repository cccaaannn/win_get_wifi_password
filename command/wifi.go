package command

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetSavedWifiNames() ([]string, error) {
	out, err := exec.Command("cmd", "/C", "netsh wlan show profile").Output()
	if err != nil {
		return nil, fmt.Errorf("can not get wifi names: %w", err)
	}

	lines := strings.Split(string(out), "All User Profile     : ")
	if len(lines) <= 1 {
		return nil, errors.New("no wifi names found")
	}

	var wifiNames []string
	for _, line := range lines[1:] {
		wifiName := strings.TrimSpace(line)
		wifiNames = append(wifiNames, wifiName)
	}

	return wifiNames, nil
}

func GetWifiPassword(wifiNetworkName string) (string, error) {
	command := fmt.Sprintf("netsh wlan show profile name=\"%s\" key=clear", wifiNetworkName)
	rawOutput, err := exec.Command("cmd", "/C", command).Output()

	if err != nil {
		return "", fmt.Errorf("can not get password: %w", err)
	}

	trimmedOutput := strings.Split(string(rawOutput), "Key Content            : ")
	if len(trimmedOutput) != 2 {
		return "", errors.New("can not get password")
	}

	splitList := strings.Split(string(trimmedOutput[1]), "\n")
	if len(splitList) == 0 {
		return "", errors.New("can not get password")
	}

	password := strings.TrimSpace(splitList[0])

	return password, nil
}
