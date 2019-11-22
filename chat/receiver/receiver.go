package receiver

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	kafkaconfig "github.com/reeechart/foroom/config"
)

type Receiver struct {
	Consumer         sarama.Consumer
	InterruptChannel chan os.Signal
}

func NewReceiver(interruptChan chan os.Signal) Receiver {
	return Receiver{
		Consumer:         getConsumer(),
		InterruptChannel: interruptChan,
	}
}

func (receiver Receiver) ConsumeMessages(topic string, partition int32, offset int64) {
	consumer, err := receiver.Consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		panic(err)
	}
	defer receiver.closeConsumer()

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			fmt.Println(msg.Value)
		case <-receiver.InterruptChannel:
			break
		}
	}
}

func (receiver Receiver) closeConsumer() {
	err := receiver.Consumer.Close()
	if err != nil {
		panic(err)
	}
}

func getConsumer() sarama.Consumer {
	brokers := getBrokersList()
	config := getSaramaConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	return consumer
}

func getBrokersList() []string {
	return []string{kafkaconfig.GetBrokerURL()}
}

func getSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	return config
}
