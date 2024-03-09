package main

import (
	"fmt"
	"os"
	"win_get_wifi_password/command"
	"win_get_wifi_password/util"
)

func main() {

	fmt.Print("Win wifi password getter\n\n")

	wifiNames, err := command.GetSavedWifiNames()
	if err != nil {
		fmt.Printf("Failed to get wifi names: %v\n", err)
		os.Exit(1)
	}

	// Menu
	fmt.Printf("Select a wifi network to get password\n\n")
	fmt.Printf("0- Exit program\n")
	for index, wifiName := range wifiNames {
		fmt.Printf("%d- %s\n", index+1, wifiName)
	}

	selectedIndex, err := util.GetUserInputAsInt()

	if err != nil || selectedIndex < 0 || selectedIndex > len(wifiNames) {
		fmt.Println("Invalid input. Please enter a number corresponding to a wifi network.")
		os.Exit(1)
	}

	if selectedIndex == 0 {
		os.Exit(0)
	}

	// Decrement the selected index to match the array index
	selectedIndex--

	wifiName := wifiNames[selectedIndex]
	password, err := command.GetWifiPassword(wifiName)
	if err != nil {
		fmt.Println("Can not get wifi password")
		os.Exit(1)
	}

	fmt.Printf("\nWifi Name: %q\nPassword: %q\n", wifiName, password)
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
