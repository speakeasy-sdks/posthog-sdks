package .models.operations;

import .utils.SpeakeasyMetadata;
public class PostTrackRequest {
    @SpeakeasyMetadata("request:mediaType=application/json")
    public Object request;
    public PostTrackRequest withRequest(Object request) {
        this.request = request;
        return this;
    }
}