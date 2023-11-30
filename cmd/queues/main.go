// Тут работаем с очередями.
//
// Когда мы отправляем сообщение в очередь, оно
// попадает только одному подписчику из группы. НАТС как бы
// распределяет нагрузку
//
// Вывод из этого упражнения такой: надо создавать
// буферезированные каналы, потому что НАТС ждать не будет.
// он тупо будет дропать сообщение, потому что
// consumer too slow
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	// закажем контекст
	ctx, cancel := context.WithCancel(context.Background())
	// сделаем группу ожидания
	started := sync.WaitGroup{}
	finish := sync.WaitGroup{}

	// темка такая, что после того, как завершится контекст,
	// мы отпишемся и закроем канал, когда внутренняя горутина завершится
	// мы уменьшим счетчик в вейтгруппе, и приложения закончится,
	// когда вейтгруппа дойдет до нуля

	// сделаем подписчиков в группе
	for i := 0; i < 5; i++ {
		i := i
		finish.Add(1)
		started.Add(1)
		go func() {
			// для каждого консьюмера персональное соединение
			nc, err := nats.Connect("nats")
			if err != nil {
				log.Fatal(err)
			}

			ch := make(chan *nats.Msg, 100)

			sub, err := nc.ChanQueueSubscribe("foo", "group", ch)
			if err != nil {
				log.Fatal(err)
			}

			// сделаем читателя
			go func() {
				started.Done()
				for m := range ch {
					// типа делаем какую-то работку
					time.Sleep(time.Millisecond * 100)
					// работка завершена, отчитаемся в чатик
					fmt.Printf("reader %d %s\n", i, m.Data)
				}
				fmt.Printf("reader inner gouroutine stopped %d\n", i)
				finish.Done()
			}()

			<-ctx.Done()
			sub.Drain()
			err = sub.Unsubscribe()
			if err != nil {
				log.Println(err)
			}
			close(ch)
			fmt.Printf("Reader %d stopped\n", i)
		}()
	}

	// подождем чтобы все ридеры запустились
	started.Wait()

	nc, err := nats.Connect("nats")
	if err != nil {
		log.Fatal(err)
	}

	// теперь запишем в эти каналы
	for i := 0; i < 100; i++ {
		bb := fmt.Sprintf("msg %d", i)
		err := nc.Publish("foo", []byte(bb))
		if err != nil {
			log.Fatal(err)
		}
	}

	// даем им какое-то время на обработку
	time.Sleep(2 * time.Second)

	// и завершаем
	cancel()

	// ждем закрытия каналов
	finish.Wait()
	fmt.Println("работа завершена")
}
