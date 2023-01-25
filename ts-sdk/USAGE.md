<!-- Start SDK Example Usage -->
```typescript
import { SDK, withSecurity} from "openapi";
import { PostTrackRequest, PostTrackResponse } from "openapi/src/sdk/models/operations";
import { AxiosError } from "axios";


const sdk = new SDK();
    
const req: PostTrackRequest = {
  request: "sit",
};

sdk.postTrack(req).then((res: PostTrackResponse | AxiosError) => {
   // handle response
});
```
<!-- End SDK Example Usage -->