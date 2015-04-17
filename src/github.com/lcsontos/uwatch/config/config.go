//
// Copyright (C) 2015-present  Laszlo Csontos
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>. */
//

package config

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

const _CONFIG_XML = "config.xml"

type (
	Parameter struct {
		XMLName xml.Name `xml:"parameter"`
		Name    string   `xml:"name"`
		Value   string   `xml:"value"`
	}

	Config struct {
		XMLName    xml.Name    `xml:"config"`
		Parameters []Parameter `xml:"parameter"`
	}
)

var parameterMap = make(map[string]string)
var parameterMutex = &sync.RWMutex{}

func GetValue(name string) string {
	defer parameterMutex.RUnlock()

	parameterMutex.RLock()

	return parameterMap[name]
}

func Init(configFile ...string) {
	cfgFile := _CONFIG_XML

	if (configFile != nil) && len(configFile) > 0 {
		cfgFile = configFile[0]
	}

	cfgFile, err := filepath.Abs(cfgFile)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reading configuration: %s", cfgFile)

	config := ReadConfig(cfgFile)

	log.Printf("Initializing application with configuration:\n%s", config)

	for _, parameter := range config.Parameters {
		SetValue(parameter.Name, parameter.Value)
	}
}

func ReadConfig(configFile string) *Config {
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
	}

	return parseConfig(data)
}

func SetValue(name, value string) {
	defer parameterMutex.Unlock()

	parameterMutex.Lock()

	parameterMap[name] = value
}

func (config *Config) String() string {
	sb := bytes.NewBufferString("")

	for _, parameter := range config.Parameters {
		sb.WriteString(parameter.Name)
		sb.WriteString("=")
		sb.WriteString(parameter.Value)
	}

	return sb.String()
}

func parseConfig(data []byte) *Config {
	config := &Config{}

	err := xml.Unmarshal(data, config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
