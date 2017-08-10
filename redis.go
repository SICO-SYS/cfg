/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cfg

type RedisInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Auth string `json:"auth"`
}

type Redis struct {
	Public RedisInfo `json:"public"`
}
