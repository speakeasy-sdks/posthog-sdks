package .models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
public class GeneralEvent {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("asdfzxcv")
    public String asdfzxcv;
    public GeneralEvent withAsdfzxcv(String asdfzxcv) {
        this.asdfzxcv = asdfzxcv;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("distinct_id")
    public String distinctId;
    public GeneralEvent withDistinctId(String distinctId) {
        this.distinctId = distinctId;
        return this;
    }
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("event")
    public String event;
    public GeneralEvent withEvent(String event) {
        this.event = event;
        return this;
    }
}