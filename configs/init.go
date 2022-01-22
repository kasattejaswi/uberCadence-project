package configs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kasattejaswi/uberCadence-project/statics"
)

func WriteConfigFile(path string, isForce bool) {
	pathToWrite, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintln("Error getting user's home directory: ", err))
	}
	if path != "" {
		pathToWrite = path
	}
	if path == "." {
		p, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintln("Error getting currenct directory path:", err))
		}
		pathToWrite = p
	}
	if isForce {
		fmt.Println("Taking backup of existing configuration if any")
		backupFileName := statics.ConfigFileName + "." + strconv.Itoa(int(time.Now().Unix()))
		buff, err := ioutil.ReadFile(filepath.Join(pathToWrite, statics.ConfigFileName))
		if err != nil {
			fmt.Println("No existing configuration found: ", err)
		}
		err = os.WriteFile(filepath.Join(pathToWrite, backupFileName), buff, 0755)
		if err != nil {
			panic(fmt.Sprintln("Unable to create backup of existing configuration file: ", err))
		}
		fmt.Println("Backup created successfully")
	} else {
		isFilePresent := CheckFileExistence(filepath.Join(pathToWrite, statics.ConfigFileName))
		if isFilePresent {
			panic("Configuration file already present. Use --force to override")
		}
	}
	pathToWrite = filepath.Join(pathToWrite, statics.ConfigFileName)
	fmt.Println("Generating the default configuration file:", pathToWrite)
	buff, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		panic(fmt.Sprintln("Internal error occurred: ", err))
	}
	err = os.WriteFile(pathToWrite, buff, 0755)
	if err != nil {
		panic(fmt.Sprintln("Unable to generate configuration file: ", err))
	}
	fmt.Println("Configuration generated successfully. Edit it accordingly")
}

func CheckFileExistence(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(fmt.Sprintln("Internal error. Unable to detect existence of configuration: ", err))
	}
}
