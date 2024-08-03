package model

type CaptureRequest struct {
	Urgency         int
	SuccessCallback func(imageData string)
	FailureCallback func(errorMessage string)
}
