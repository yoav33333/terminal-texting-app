package network

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stefanwichmann/lanscan"
	"log"
	"net/url"
	"strings"
	"textEditor/util"
	"time"
)

const port = 3000

func ping(ip string) (bool, string) {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%v", ip, port), Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	//defer c.Close()

	if err != nil {
		log.Fatal("dial:", err)
	}
	err = c.WriteMessage(websocket.TextMessage, []byte("ping"))
	if err != nil {
		return false, ""
	}
	_, message, err := c.ReadMessage()
	c.Close()
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
func checkIfNameExists(name string) bool {
	for _, user := range util.Users {
		if user == name {
			return true
		}
	}
	return false
}

func RunNetworkingShit() {
	go StartWebSocketServer()
	for {
		ips := findDevicesInNetwork()
		util.Users = make(map[string]string)
		for _, ip := range ips {
			active, userName := ping(ip)
			if active && !checkIfNameExists(userName) {
				println("Client: Found WebSocket server at:", ip)
				util.Users[ip] = userName
			}
		}
		for ip, user := range util.Users {
			println("Client: User", user, "is at", ip)
		}
		time.Sleep(5 * time.Second)
	}
}
