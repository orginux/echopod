package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

const templ = `Hostname: {{ .Hostname}}`

var (
	hostname string
	ipAddr   string
)

func main() {
	type PodInfo struct {
		Hostname string
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error getting hostname:", err)
	}
	info := PodInfo{hostname}

	report := template.Must(template.New("podinfo").Parse(templ))

	if err := report.Execute(os.Stdout, info); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", hostname, r.RemoteAddr, r.Method, r.URL)
		report.Execute(w, info)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

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
