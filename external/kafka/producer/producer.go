package producer

type Producer struct {
	brokers []string
}

type ProducerInterface interface {
	Produce(topic string, message []byte) error
}

func NewProducer(brokers []string) ProducerInterface {
	return &Producer{brokers: brokers}
}

func (c *Producer) Produce(topic string, message []byte) error {
	//publish messages to kafka
	return nil
}
