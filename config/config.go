package config

const (
	KAFKA_HOST = "localhost"
	KAFKA_PORT = "9092"
)

type KafkaConfig struct {
	Host string
	Port string
}

func (config KafkaConfig) GetBrokerURL() string {
	return config.Host + ":" + config.Port
}

func GetKafkaBroker() KafkaConfig {
	return KafkaConfig{
		Host: KAFKA_HOST,
		Port: KAFKA_PORT,
	}
}
