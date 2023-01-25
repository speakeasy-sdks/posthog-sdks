import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class EventsCaptureResponse extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=status" })
  status?: number;
}
