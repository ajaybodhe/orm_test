package queues

import (
	"github.com/ajaybodhe/orm_test/constants"	
//	"github.com/bitly/go-nsq"
//	"log"
)

type Queue interface {
	Publish(string) bool
	Consume() string
	Stop() bool
}

type HandlerFunc func([]byte) error

func QueueCreation(queueType string, handler HandlerFunc, topic string) Queue {
	if queueType == constants.NSQ {
		var qi Queue
		qi = &NsqQueue{Handler: handler, Topic:topic}
		return qi
	}
	return nil
}