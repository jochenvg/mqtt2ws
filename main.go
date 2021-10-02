package main

import (
	"flag"
	"io"
	"log"
	"net"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/eclipse/paho.mqtt.golang/packets"
)

var listen = flag.String("listen", ":1884", "listen tcp socket")
var url = flag.String("url", "ws://broker.hivemq.com:8000/mqtt", "websocket url")

func main() {
	// Parse Flags
	flag.Parse()

	// Listen on TCP Socket
	log.Println("Listening on:", *listen)
	listener, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Println("Listen error", err)
		return
	}
	defer listener.Close()

	// Set up context with cancel

	// Loop forever
	for {
		// Accept TCP connection
		s, err := listener.Accept()
		if err != nil {
			log.Println("Accept error", err)
			return
		}
		go func() {
			// Defer closing TCP connection
			defer s.Close()

			// Create MQTT-WS client
			ws, err := paho.NewWebsocket(*url, nil, 0, nil, nil)
			if err != nil {
				log.Println("Websocket error", err)
				return
			}
			// Defer closing Websocket connection
			defer ws.Close()

			// If either direction is done, both are done
			done := make(chan bool)
			// Read from s, write to ws
			go func() {
				defer func() { done <- true }()
				for {
					p, err := packets.ReadPacket(s)
					if err != nil {
						log.Println("Error reading from TCP socket", err)
						return
					}
					err = p.Write(ws)
					if err != nil {
						log.Println("Error writing to Websocket", err)
						return
					}
				}
			}()
			// Read from ws, write to s
			go func() {
				defer func() { done <- true }()
				io.Copy(s, ws)
			}()
			<-done
		}()
	}
}
