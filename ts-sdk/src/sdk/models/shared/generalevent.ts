import { SpeakeasyMetadata, SpeakeasyBase } from "../../../internal/utils";



export class GeneralEvent extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "json, name=asdfzxcv" })
  asdfzxcv?: string;

  @SpeakeasyMetadata({ data: "json, name=distinct_id" })
  distinctId?: string;

  @SpeakeasyMetadata({ data: "json, name=event" })
  event?: string;
}
