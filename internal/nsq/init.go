package nsq

import (
	"fmt"

	"code.meikeland.com/me/forwarding/internal/conf"
	"github.com/meikeland/logger"
	"github.com/nsqio/go-nsq"
)

var (
	// producer NSQ生产者
	producer *nsq.Producer
	// producer NSQ消费者
	consumer *nsq.Consumer
	// nsqHost NSQ连接地址
	nsqHost string
)

// Init 初始化
func Init() {
	nsqConf := conf.GetNSQ()
	nsqHost = nsqConf
	initProducer()
}

// initProducer 初始化生产者
func initProducer() {
	fmt.Print(nsqHost)
	var err error
	producer, err = nsq.NewProducer(nsqHost, nsq.NewConfig())
	if err != nil {
		logger.Panicf("init nsq by order close error : %s", err.Error())
		panic(fmt.Sprintf("init nsq by order close error : %s", err.Error()))
	}
}

// initConsumer 初始化消费者
func initConsumer(handler nsq.Handler, topic, channel string) {
	var err error
	consumer, err = nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		logger.Panicf("init consumer by order close error : %s", err.Error())
		panic(fmt.Sprintf("init consumer by order close error : %s", err.Error()))
	}
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQD(nsqHost)
	if err != nil {
		logger.Panicf("init consumer connect nsq by order close error : %s", err.Error())
		panic(fmt.Sprintf("init consumer connect nsq by order close error : %s", err.Error()))
	}
}
