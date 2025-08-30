package main

import "textEditor/network"

func main() {
	network.RunNetworkingShit()
	//RunTexting()
}

//func main() {
//	// Scan for hosts listening on tcp port 80.
//	// Use 20 threads and timeout after 5 seconds.
//	hosts := lanscan.LinkLocalAddresses("ip4")
//
//	for i, host := range hosts {
//		hosts[i] = strings.Split(host, "/")[0]
//	}
//	for _, host := range hosts {
//		println(host)
//	}
//}
