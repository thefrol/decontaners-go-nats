// Этот пример показывает как асинхронно запускать процендуры
// у подписчика при помощи request/response
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		_, err := nc.Subscribe("foo", func(msg *nats.Msg) {
			//doing stuff for 2s
			fmt.Println("*Subscriber doing work for 2s*")
			time.Sleep(2 * time.Second)
			msg.Respond([]byte("this is a response"))
		})
		if err != nil {
			log.Fatal(err)
		}

	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Sending")
	msg, err := nc.Request("foo", []byte("le request"), time.Second*10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("message received: ", string(msg.Data))

	// sub.Drain()
}
