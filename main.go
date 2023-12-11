package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const DefaultDirectory = "./files"

func main() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	var ip []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = append(ip, ipnet.IP.String())
			}
		}
	}

	dir := DefaultDirectory
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Serving files from '%s' directory.", dir)
	fileServer := http.FileServer(http.Dir(dir))

	// Set the route for the file server.
	http.Handle("/", fileServer)

	for port := 8080; port <= 8100; port++ {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			if strings.Contains(err.Error(), "bind: address already in use") {
				log.Printf("Port %d is in use", port)
				continue
			} else {
				log.Fatal(err)
			}
		}

		for _, ip := range ip {
			log.Printf("Server started on http://%s:%d", ip, port)
		}

		http.Serve(listener, nil)

		break
	}
}
