<!-- Start SDK Example Usage -->
```java
package hello.world;

import .SDK;
import .models.shared.Security;

public class Application {
    public static void main(String[] args) {
        try {
            SDK.Builder builder = SDK.builder();

            SDK sdk = builder.build();

            PostTrackRequest req = new PostTrackRequest() {{
                request = "sit";
            }};

            PostTrackResponse res = sdk.postTrack(req);

            if (res.eventsCaptureResponse.isPresent()) {
                // handle response
            }
        } catch (Exception e) {
            // handle exception
        }
```
<!-- End SDK Example Usage -->