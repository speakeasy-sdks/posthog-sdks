package .models.operations;


public class PostTrackResponse {
    public String contentType;
    public PostTrackResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public .models.shared.EventsCaptureResponse eventsCaptureResponse;
    public PostTrackResponse withEventsCaptureResponse(.models.shared.EventsCaptureResponse eventsCaptureResponse) {
        this.eventsCaptureResponse = eventsCaptureResponse;
        return this;
    }
    public Long statusCode;
    public PostTrackResponse withStatusCode(Long statusCode) {
        this.statusCode = statusCode;
        return this;
    }
}