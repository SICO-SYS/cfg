/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

Contributors

*/

package cfg

import (
	"encoding/json"
	"io/ioutil"
)

func ReadLocalFile() []byte {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil
	}
	return data
}

func Map2struct(m map[string]string, c *ConfigItems) {
	data, _ := json.Marshal(m)
	err := json.Unmarshal(data, c)
	if err != nil {
		c = &ConfigItems{}
	}
}

func Unmarshal(data []byte, c *ConfigItems) {
	err := json.Unmarshal(data, c)
	if err != nil {
		c = &ConfigItems{}
	}
}
