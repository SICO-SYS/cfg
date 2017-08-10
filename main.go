/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Config *ConfigItem
)

type ConfigItem struct {
	Version    string
	AAAEnable  bool       `json:"AAAEnable"`
	OpenAccess OpenAccess `json:"OpenAccess"`
	Redis      Redis      `json:"redis"`
	Mysql      Mysql      `json:"mysql"`
	Mongo      Mongo      `json:"mongo"`
	Rabbit     Rabbit     `json:"rabbit"`
	RPCPort    RPCPort    `json:"rpcport"`
	Sentry     Sentry     `json:"sentry"`
}

func (c *ConfigItem) ReadConfig(path string) (*ConfigItem, error) {

	// We'll change to Call configAPI to parse the config
	// ioutil.ReadFile("https://cfg.siner.io/production")
	d, err := ioutil.ReadFile(path)
	if err == nil {
		err = json.Unmarshal(d, &c)
	}

	if c.Version == "" {
		c.Version = "0.2.3"
	}

	return c, err
}

func init() {
	defer func() {
		if emsg := recover(); emsg != nil {
			log.Println(emsg)
		}
	}()
	var err error
	Config, err = (&ConfigItem{}).ReadConfig("config.json")
	if err != nil {
		log.Println(err)
		log.Fatalln("Parse config.json problem, make sure file with correct syntax")
	}
}
