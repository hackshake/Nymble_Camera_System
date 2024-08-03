package service

import "nymble_camera_system/internal/model"

type CameraSystemService struct {
	rm RequestManagerInterface
}

type CameraSystemServiceInterface interface {
	SubmitCaptureRequest(request *model.CaptureRequest)
}

func NewCameraSystemService(rm RequestManagerInterface) CameraSystemServiceInterface {
	return &CameraSystemService{rm}
}

func (css *CameraSystemService) SubmitCaptureRequest(request *model.CaptureRequest) {
	css.rm.AddRequest(request)
}
