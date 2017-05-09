/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cfg

type MongoInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Addr string `json:"addr"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Mongo struct {
	Default MongoInfo `json:"default"`
	User    MongoInfo `json:"user"`
	Cloud   MongoInfo `json:"user"`
	Asset   MongoInfo `json:"user"`
	Geo     MongoInfo `json:"geo"`
}
