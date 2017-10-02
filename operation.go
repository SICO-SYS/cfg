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

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/dao/grpc"
)

func ReadLocalFile() []byte {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil
	}
	return data
}

func ReadFilePath(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadConfigServer() []byte {
	cc, err := rpc.Conn("boron", "6666")
	if err != nil {
		return nil
	}
	defer cc.Close()
	r := rpc.ConfigPullRPC(cc, &pb.ConfigPullCall{Id: "system", Environment: "config"})
	return r.Params
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
