package main

import (
	"bufio"
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
	namespace, err := getNamespace()
	if err != nil {
		log.Fatal("Error getting Pod namespace:", err)
	}

	webServer(hostname, IPaddr, namespace)
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

func getNamespace() (string, error) {
	var namespace string
	namespaceFile := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	if _, err := os.Stat(namespaceFile); err == nil {
		file, err := os.Open(namespaceFile)
		if err != nil {
			return "", err
		}
		defer file.Close()

		scaner := bufio.NewScanner(file)
		for scaner.Scan() {
			namespace += scaner.Text()
		}
		return namespace, nil
	}
	return "", nil
}

func webServer(hostname, IPaddr, namespace string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", hostname, r.RemoteAddr, r.Method, r.URL)
		fmt.Fprintf(w, "Name: %s\nIP: %s\nNamespace: %s\nURI: %s\n", hostname, IPaddr, namespace, r.RequestURI)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
