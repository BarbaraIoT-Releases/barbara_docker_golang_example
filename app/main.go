package main

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"os"
	"time"
)

const (
	appconfigFile = "/appconfig/appconfig.json"
	appSleepTime = time.Duration(5)*time.Second
)

var (
	config1 = "value1"
	config2 = "value2"
)

func main() {
	for {
		if fileExists(appconfigFile){
			fmt.Println("Appconfig file found")
			appConfigParsed, _ := parseJSON(appconfigFile)
			tmpConfig1 := appConfigParsed.S("config1").Data()
			if tmpConfig1 != nil {
				config1 = fmt.Sprintf("%v", tmpConfig1)
			}
			tmpConfig2 := appConfigParsed.S("config2").Data()
			if tmpConfig2 != nil {
				config2 = fmt.Sprintf("%v", tmpConfig2)
			}
		} else {
			fmt.Println("No appconfig file exists, dafault values used")
		}

		fmt.Println("App config values:")
		fmt.Println(" - config 1: "+config1)
		fmt.Println(" - config 2: "+config2)

		time.Sleep(appSleepTime)
	}
}

func fileExists(filePath string) bool {
	_, statErr := os.Stat(filePath)
	if os.IsNotExist(statErr) {
		return false
	}

	return true
}

func parseJSON(jsonPath string) (*gabs.Container, error) {
	jsonFile, jsonFileErr := os.Open(jsonPath)
	if jsonFileErr != nil {
		return nil, errors.New("parseJSON: " + jsonFileErr.Error())
	}
	defer func() {
		jsonFileErr := jsonFile.Close()
		if jsonFileErr != nil {
			return
		}
	}()
	jsonBytes, jsonBytesErr := ioutil.ReadAll(jsonFile)
	if jsonBytesErr != nil {
		return nil, errors.New("failed to parse json file: " + jsonPath + ". Error: " + jsonBytesErr.Error())
	}

	return gabs.ParseJSON(jsonBytes)
}
