package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	model "nymble_camera_system/internal/model"
	service "nymble_camera_system/internal/service"
	"sync"
)

type CameraSystemHandler struct {
	service service.CameraSystemServiceInterface
}

type CameraSystemHandlerInterface interface {
	CaptureImageHandler() http.HandlerFunc
}

func NewCameraSystemHandler(service service.CameraSystemServiceInterface) CameraSystemHandlerInterface {
	return &CameraSystemHandler{
		service: service,
	}
}

func (csh *CameraSystemHandler) CaptureImageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request model.CaptureRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Define the callbacks
		var responseLock sync.Mutex
		var apiResponse string
		var apiError error

		successCallback := func(result string) {
			responseLock.Lock()
			apiResponse = result
			apiError = nil
			responseLock.Unlock()
		}

		failureCallback := func(errorMessage string) {
			responseLock.Lock()
			apiResponse = errorMessage
			apiError = fmt.Errorf(errorMessage)
			responseLock.Unlock()
		}

		request.SuccessCallback = successCallback
		request.FailureCallback = failureCallback

		csh.service.SubmitCaptureRequest(&request)

		//This ensures that the HTTP response is only sent after the callbacks have finished processing.
		responseLock.Lock()
		defer responseLock.Unlock()

		if apiError != nil {
			http.Error(w, apiResponse, http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(apiResponse))
		}
	}
}
