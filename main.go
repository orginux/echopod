package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	hostname string
	IPaddr   string
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error getting hostname:", err)
	}

	IPaddr, err := getIP()
	if err != nil {
		log.Fatal("Error getting IP address:", err)
	}

	webServer(hostname, IPaddr)
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("Error getting Interface Addrs:\n %v", err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}

func webServer(hostname, IPaddr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", hostname, r.RemoteAddr, r.Method, r.URL)
		fmt.Fprintf(w, "Hostname: %s\nIP: %s\nURI: %s\n", hostname, IPaddr, r.RequestURI)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
