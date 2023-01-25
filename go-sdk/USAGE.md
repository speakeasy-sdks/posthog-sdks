<!-- Start SDK Example Usage -->
```go
package main

import (
    "openapi"
    "openapi/pkg/models/shared"
    "openapi/pkg/models/operations"
)

func main() {
    s := posthog.New()
    
    req := operations.PostTrackRequest{
        Request: "sit",
    }
    
    res, err := s.PostTrack(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventsCaptureResponse != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->