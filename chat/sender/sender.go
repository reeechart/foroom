package sender

import (
	"github.com/Shopify/sarama"
	kafkaconfig "github.com/reeechart/foroom/config"
)

type Sender struct {
	Producer sarama.SyncProducer
}

func NewSender() Sender {
	return Sender{
		Producer: getProducer(),
	}
}

func (sender Sender) SendMessage(topic string, content string) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(content),
	}
	partition, offset, err = sender.Producer.SendMessage(msg)
	return
}

func getProducer() sarama.SyncProducer {
	brokers := getBrokersList()
	config := getSaramaConfig()
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	return producer
}

func getBrokersList() []string {
	return []string{kafkaconfig.GetBrokerURL()}
}

func getSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	return config
}
