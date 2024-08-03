package main

import (
	"fmt"
	"log"
	"net/http"
	consumer "nymble_camera_system/external/kafka/consumer"
	"nymble_camera_system/external/kafka/producer"
	handler "nymble_camera_system/internal/handler"
	service "nymble_camera_system/internal/service"
)

func main() {
	startConsumer()
	startServer()
}

func startServer() {
	cameraController := &service.CameraController{}

	// Registering the callbacks to the camera controller
	cameraController.RegisterSuccessCallback(func(imageData string) {
		fmt.Printf("Success: %s\n", imageData)
	})

	cameraController.RegisterFailureCallback(func(errorMessage string) {
		fmt.Printf("Failure: %s\n", errorMessage)
	})

	producer := producer.NewProducer([]string{"localhost:9092"})
	topics := map[int]string{
		0: "low_priority",
		1: "medium_priority",
		2: "high_priority",
	}
	rm := service.NewRequestManager(producer, topics)
	cameraService := service.NewCameraSystemService(rm)
	cameraHandler := handler.NewCameraSystemHandler(cameraService)
	http.HandleFunc("/capture", cameraHandler.CaptureImageHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startConsumer() {
	consumer.Start()
}
