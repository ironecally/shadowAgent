package shadowAgent

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go NewQueue(&Option{
		TotalWorker: 5,
	})

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()

	for {
		// time.Sleep(1 * time.Second)
		Enqueue(context.Background(), Job{
			JobName: "test1",
			HandlerFunc: func(message []byte) error {
				log.Println(string(message))
				return nil
			},
			Message: []byte("hello 世界!"),
		})
		// Enqueue(context.Background(), Job{
		// 	JobName: "test2",
		// 	HandlerFunc: func(message []byte) error {
		// 		log.Println(string(message))
		// 		return nil
		// 	},
		// 	Message: []byte("hello world!"),
		// })
	}

}
