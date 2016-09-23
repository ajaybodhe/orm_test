package queues

import (
	"log"
	"github.com/bitly/go-nsq"
	"github.com/ajaybodhe/orm_test/conf"	
)

type NsqQueue struct {
	nsqProd *nsq.Producer
	nsqCons *nsq.Consumer
	Handler HandlerFunc
	Topic string
}

func (this * NsqQueue) HandleMessage(m * nsq.Message) error {
	if this.Handler != nil {
		err:=this.Handler(m.Body)
		return err
	}
	return nil
}

func (this * NsqQueue) createPublisher() {
	var err error
	config := nsq.NewConfig()
	this.nsqProd, err = nsq.NewProducer(conf.OrmTestConfig.NSQ.Host + ":" + conf.OrmTestConfig.NSQ.Port, config)
	
	if err != nil {
		log.Panic("can not create nsq producer")
	}
}

func (this *NsqQueue) Publish(msg string) bool{
	if this.nsqProd == nil {
		this.createPublisher()
	}
	err := this.nsqProd.Publish(this.Topic, []byte(msg))
	if err != nil {
	    log.Panic("Could not connect and publish to nsq")
		return false
	}	
	return true
}

func (this *NsqQueue) createConsumer() error{
	var err error
	
	config := nsq.NewConfig()
	
	this.nsqCons, err = nsq.NewConsumer(this.Topic, conf.OrmTestConfig.Consumer.ChannelName, config)
	if err != nil {
		log.Panic("can not create nsq consumer")
		return err
	}
	
	this.nsqCons.AddHandler(this)
	
	err = this.nsqCons.ConnectToNSQD(conf.OrmTestConfig.NSQ.Host +":" + conf.OrmTestConfig.NSQ.Port)
	if err != nil {
		log.Panic("Could not connect to nsq")
		return err
  	}
	return nil
}


func (this *NsqQueue) Consume() string{
	this.createConsumer()
	return ""
}
//func (this *NsqQueue) Consume(this.Topic string, msg string) {
//	if this.nsqCons == nil {
//		this.createConsumer()
//	}
//	err := this.nsqProd.Publish(this.Topic, []byte(msg))
//	if err != nil {
//	    log.Panic("Could not connect and publish to nsq")
//	}	
//}

func (this *NsqQueue) Stop() bool {
	if this.nsqProd != nil {
		this.nsqProd.Stop()
	}
	/* TBD Ajay how do we stop consumer */
//	if this.nsqCons != nil {
//		this.nsqCons.Stop()
//	}
	return true
}
