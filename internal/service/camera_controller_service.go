package service

import (
	"math/rand"
	model "nymble_camera_system/internal/model"
	"time"
)

type CameraController struct {
	successCallback func(imageData string)
	failureCallback func(errorMessage string)
}

func (cc *CameraController) StartCapture(request *model.CaptureRequest) {
	go func() {
		time.Sleep(time.Second)   // Simulating capture delay
		if rand.Float32() < 0.9 { // Simulating 90% success rate
			if cc.successCallback != nil {
				cc.successCallback("captured_image_data")
			}
		} else {
			if cc.failureCallback != nil {
				cc.failureCallback("camera_error")
			}
		}
	}()
}

func (cc *CameraController) RegisterSuccessCallback(callback func(imageData string)) {
	cc.successCallback = callback
}

func (cc *CameraController) RegisterFailureCallback(callback func(errorMessage string)) {
	cc.failureCallback = callback
}
