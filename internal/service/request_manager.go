package service

import (
	"nymble_camera_system/external/kafka/producer"
	"nymble_camera_system/internal/model"
)

type RequestManager struct {
	producer producer.ProducerInterface
	topics   map[int]string // Mapping urgency to Kafka topics

}

type RequestManagerInterface interface {
	AddRequest(request *model.CaptureRequest)
}

func NewRequestManager(producer producer.ProducerInterface, topics map[int]string) RequestManagerInterface {
	return &RequestManager{producer, topics}
}

func (rm *RequestManager) AddRequest(request *model.CaptureRequest) {
	topic := rm.topics[request.Urgency]
	message := serializeRequest(request) // Serialize the request
	rm.producer.Produce(topic, message)
}

func serializeRequest(request *model.CaptureRequest) []byte {
	// Serialize the request object to JSON or other format
	return []byte{}
}
