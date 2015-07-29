package config

import (
	"api/consts"
	"api/log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
	"path"
)

var CIDR string

func LoadConfig() {
	curdir, _ := os.Getwd()
	configPath := path.Join(curdir, consts.CONFIGFILENAME)
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Error(err)
		panic("Unable to read the config file " + consts.CONFIGFILENAME)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(configContent))
	var conf map[string]interface{}
	if err := decoder.Decode(&conf); err != nil {
		panic("The content of " + consts.CONFIGFILENAME + " are not valid JSON contents.")
	}

	//CIDR Config load
	if _cidr, k := conf[consts.CIDR_IP]; k {
		if cidr, k := _cidr.(string); k {
			CIDR = cidr
		}
	}
	if len(CIDR) < 1 {
		panic(fmt.Errorf("NO CIDR Configuration found in config."))
	}

}
