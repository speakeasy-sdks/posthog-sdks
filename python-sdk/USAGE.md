<!-- Start SDK Example Usage -->
```python
import sdk
from sdk.models import operations, shared

s = sdk.SDK()
    
req = operations.PostTrackRequest(
    request="sit",
)
    
res = s.post_track_(req)

if res.events_capture_response is not None:
    # handle response
```
<!-- End SDK Example Usage -->