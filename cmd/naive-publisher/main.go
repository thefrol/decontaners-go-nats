package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// Конкурентный подписчик
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Получено сообщение: %s\n", string(m.Data))
	})

	// Простой издатель
	nc.Publish("foo", []byte("Hello World"))

	time.Sleep(100 * time.Millisecond)

}
