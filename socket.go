package seadisc

import "strings"

var (
	BaseUrl string = "wss://{network}stream.openseabeta.com/socket/websocket?token="
)

func GetSocketUrl(network, token string) string {
	var (
		net string = ""
		url string
	)

	if strings.ToLower(network) != "main" {
		net = "testnets-"
	}

	url = strings.Replace(BaseUrl, "{network}", net, 1)
	return url + token
}
