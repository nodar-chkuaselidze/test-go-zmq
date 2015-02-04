package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"time"
)

const XREQ = "tcp://localhost:3002"

func main() {
	worker, _ := zmq.NewSocket(zmq.REP)

	set_id(worker)
	defer worker.Close()

	worker.Connect(XREQ)

	identity, _ := worker.GetIdentity()
	fmt.Println("my identity is: ", identity)

	for {
		request, _ := worker.RecvMessage(0)

		fmt.Println("Got request: ", request)
		worker.SendMessage("Worker do your job")
		fmt.Println(".. next")
	}
}

func set_id(soc *zmq.Socket) {
	rand.Seed(time.Now().UnixNano())
	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	soc.SetIdentity(identity)
}
