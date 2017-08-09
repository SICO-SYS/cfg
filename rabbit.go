/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cfg

type RabbitInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"user"`
	Password string `json:"pass"`
	Topic    string `json:"topic"`
}

type Rabbit struct {
	Testing RabbitInfo `json:"testing"`
}
