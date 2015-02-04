package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"time"
)

const XREP = "tcp://localhost:3001"

func main() {
	worker, _ := zmq.NewSocket(zmq.REQ)

	set_id(worker)
	defer worker.Close()

	worker.Connect(XREP)

	identity, _ := worker.GetIdentity()
	fmt.Println("my identity is: ", identity)

	total := 0
	for {
		worker.SendMessage("Hi there", 0)
		reply, _ := worker.RecvMessage(0)

		total++
		fmt.Println("message: ", reply, ", No", total)
	}
}

func set_id(soc *zmq.Socket) {
	rand.Seed(time.Now().UnixNano())
	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	soc.SetIdentity(identity)
}
