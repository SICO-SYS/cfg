/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cfg

type MongoInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Mongo struct {
	Testing MongoInfo `json:"testing"`
	User    MongoInfo `json:"user"`
	Cloud   MongoInfo `json:"cloud"`
	Asset   MongoInfo `json:"asset"`
}
