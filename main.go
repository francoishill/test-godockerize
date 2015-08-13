package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func printIP() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting IPs: " + err.Error())
		return
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("Error getting IP address " + err.Error())
			continue
		}

		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				// ip = v.IP
			case *net.IPAddr:
				ip = v.IP
				fmt.Println("Found IP " + ip.String())
			}

		}
	}
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	printIP()
	log.Fatal(http.ListenAndServe(":45678", nil))
}
