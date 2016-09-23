package main

import (
	"github.com/ajaybodhe/orm_test/queues"
	"github.com/ajaybodhe/orm_test/constants"
	"github.com/ajaybodhe/orm_test/conf"
	"sync"
	"log"
)

func FakeHandler(msg []byte) error {
	log.Printf("Got a message: %v", string(msg))
//	wg.Done()
    return nil
}
func main() {
	wg := &sync.WaitGroup{}
  	wg.Add(1)
	que := queues.QueueCreation(constants.NSQ, FakeHandler, conf.OrmTestConfig.Queue.Topic)
	que.Consume()
   	wg.Wait()
}

