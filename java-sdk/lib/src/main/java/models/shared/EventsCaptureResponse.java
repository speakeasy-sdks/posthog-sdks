package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class EventsCaptureResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("status")
    public Double status;
    public EventsCaptureResponse withStatus(Double status) {
        this.status = status;
        return this;
    }
}