package posthog

import (
	"net/http"

	"context"
	"fmt"
	"openapi/pkg/models/operations"
	"openapi/pkg/models/shared"
	"openapi/pkg/utils"
	"strings"
)

var ServerList = []string{
	"https://posthog.com/",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Posthog struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*Posthog)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Posthog) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Posthog) {
		sdk._defaultClient = client
	}
}

func New(opts ...SDKOption) *Posthog {
	sdk := &Posthog{
		_language:   "go",
		_sdkVersion: "",
		_genVersion: "0.21.0",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk._defaultClient == nil {
		sdk._defaultClient = http.DefaultClient
	}
	if sdk._securityClient == nil {

		sdk._securityClient = sdk._defaultClient

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	return sdk
}

// PostTrack - Capture an event
// Capture an event. Events are the core of PostHog, and are what you use
// to track user behavior, and then analyze and visualize in PostHog.
func (s *Posthog) PostTrack(ctx context.Context, request operations.PostTrackRequest) (*operations.PostTrackResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/track/"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s._defaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostTrackResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EventsCaptureResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EventsCaptureResponse = out
		}
	case httpRes.StatusCode == 400:
	}

	return res, nil
}
