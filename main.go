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
	Config *Main
	err    error
)

type Main struct {
	Version    string     `json:"version"`
	Redis      Redis      `json:"redis"`
	Mysql      Mysql      `json:"mysql"`
	Mongo      Mongo      `json:"mongo"`
	Rabbit     Rabbit     `json:"rabbit"`
	RPC        RpcPort    `json:"rpc"`
	Sentry     Sentry     `json:"sentry"`
	OpenAccess OpenAccess `json:"OpenAccess"`
}

func (c *Main) ReadConfig(path string) (*Main, error) {

	// We'll change to Call configAPI to parse the config
	// ioutil.ReadFile("https://cfg.siner.io/production")
	d, err := ioutil.ReadFile(path)
	if err == nil {
		err = json.Unmarshal(d, &c)
	}

	return c, err
}

func init() {
	defer func() {
		if emsg := recover(); emsg != nil {
			log.Println(emsg)
		}
	}()
	Config, err = (&Main{}).ReadConfig("config.json")
	if err != nil {
		log.Println(err)
		log.Fatalln("Parse config.json problem, make sure file with correct syntax")
	}
}
