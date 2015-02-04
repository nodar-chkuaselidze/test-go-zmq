package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	frontend.Bind("tcp://*:3001")
	backend.Bind("tcp://*:3002")

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

	//  Switch messages between sockets
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case frontend:
				msg, _ := s.RecvMessage(0)
				printMessage("frontend", msg)

				backend.SendMessage(msg)
			case backend:
				msg, _ := s.RecvMessage(0)
				printMessage("backend", msg)

				frontend.SendMessage(msg)
			}
		}
	}
}

func printMessage(end string, msg []string) {
	fmt.Println("from ", end, ":", msg)
}
