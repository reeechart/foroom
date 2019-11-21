package config

const (
	KAFKA_HOST = "localhost"
	KAFKA_PORT = "9092"
)

type kafkaConfig struct {
	Host string
	Port string
}

func GetBrokerURL() string {
	config := getKafkaConfig()
	return config.Host + ":" + config.Port
}

func getKafkaConfig() kafkaConfig {
	return kafkaConfig{
		Host: KAFKA_HOST,
		Port: KAFKA_PORT,
	}
}
