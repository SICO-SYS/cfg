/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cfg

type Sentry struct {
	Enable bool   `json:"enable"`
	DSN    string `json:"dsn"`
}
