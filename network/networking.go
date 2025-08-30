package network

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stefanwichmann/lanscan"
	"log"
	"net/url"
	"strings"
	"textEditor/util"
)

const port = 3000

func ping(ip string) (bool, string) {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%v", ip, port), Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	defer c.Close()

	if err != nil {
		log.Fatal("dial:", err)
	}
	err = c.WriteMessage(websocket.TextMessage, []byte("ping"))
	if err != nil {
		return false, ""
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		return false, ""
	}
	log.Printf("recv: %s", message)
	return true, string(message)
}

func findDevicesInNetwork() []string {
	ips := lanscan.LinkLocalAddresses("ip4")

	for i, host := range ips {
		ips[i] = strings.Split(host, "/")[0]
	}
	return ips
}

func RunNetworkingShit() {
	go StartWebSocketServer()
	ips := findDevicesInNetwork()
	for _, ip := range ips {
		active, userName := ping(ip)
		if active {
			println("Client: Found WebSocket server at:", ip)
			util.Users[ip] = userName
			println("Users:", util.Users)
		}
	}

}
