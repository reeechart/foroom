package sender

import (
	"os"

	"github.com/Shopify/sarama"
	kafkaconfig "github.com/reeechart/foroom/config"
)

type Sender struct {
	Producer         sarama.SyncProducer
	InterruptChannel chan os.Signal
}

func NewSender(interruptChan chan os.Signal) Sender {
	return Sender{
		Producer:         getProducer(),
		InterruptChannel: interruptChan,
	}
}

func (sender Sender) ListenAndSendUserInputs(topic string, user string) {
	msgChan := make(chan string)
	inputListener := InputListener{MsgChannel: msgChan}
	for {
		go inputListener.GetInput()
		select {
		case msg := <-inputListener.MsgChannel:
			msg = user + ": " + msg
			sender.sendMessage(topic, msg)
		case <-sender.InterruptChannel:
			sender.closeProducer()
		}
	}
}

func (sender Sender) sendMessage(topic string, content string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(content),
	}

	_, _, err := sender.Producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
}

func (sender Sender) closeProducer() {
	err := sender.Producer.Close()
	if err != nil {
		panic(err)
	}
}

func getProducer() sarama.SyncProducer {
	brokers := getBrokersList()
	config := getSenderConfig()
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	return producer
}

func getBrokersList() []string {
	return []string{kafkaconfig.GetBrokerURL()}
}

func getSenderConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	return config
}
