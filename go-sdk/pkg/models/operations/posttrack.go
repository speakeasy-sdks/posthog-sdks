package operations

import (
	"openapi/pkg/models/shared"
)

type PostTrackRequestBody1 struct {
	APIKey *string       `json:"api_key,omitempty"`
	Batch  []interface{} `json:"batch,omitempty"`
}

type PostTrackRequest struct {
	Request *interface{} `request:"mediaType=application/json"`
}

type PostTrackResponse struct {
	ContentType           string
	EventsCaptureResponse *shared.EventsCaptureResponse
	StatusCode            int64
}
