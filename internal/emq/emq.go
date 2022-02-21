package emq

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Emq struct {
	Client mqtt.Client
}

func Connect(cfg Config) mqtt.Client { //nolint:ireturn
	// TODO: temporary solution
	time.Sleep(5 * time.Second)

	opts := mqtt.NewClientOptions().AddBroker(cfg.URL).SetClientID(cfg.ClientID)

	// Set the message callback handler
	defaultHandler := func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("DefaultHandler: topic: %s msg: %s\n", msg.Topic(), msg.Payload())
	}
	opts.SetDefaultPublishHandler(defaultHandler)

	opts.SetKeepAlive(60 * time.Second)   //nolint:gomnd
	opts.SetPingTimeout(10 * time.Second) //nolint:gomnd

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c
}

func (e Emq) Subscribe(topic string, callback mqtt.MessageHandler) {
	subscribe := e.Client.Subscribe(topic, 0, callback)
	subscribe.Wait()
}

func (e Emq) Publish(topic string, data interface{}) {
	token := e.Client.Publish(topic, 1, false, data)
	token.Wait()
}
