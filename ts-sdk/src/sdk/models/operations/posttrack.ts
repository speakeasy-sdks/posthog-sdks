import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";
import * as shared from "../shared";



export class PostTrackRequestBody1 extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=api_key" })
  apiKey?: string;

  @SpeakeasyMetadata({ data: "json, name=batch" })
  batch?: any[];
}


export class PostTrackRequest extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "request, media_type=application/json" })
  request?: any;
}


export class PostTrackResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  eventsCaptureResponse?: shared.EventsCaptureResponse;

  @SpeakeasyMetadata()
  statusCode: number;
}
