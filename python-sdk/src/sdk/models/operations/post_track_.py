import dataclasses
from typing import Any,Optional
from dataclasses_json import dataclass_json
from sdk import utils
from ..shared import generalevent as shared_generalevent
from ..shared import eventscaptureresponse as shared_eventscaptureresponse


@dataclass_json
@dataclasses.dataclass
class PostTrackRequestBody1:
    api_key: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('api_key') }})
    batch: Optional[list[Any]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.field_name('batch') }})
    

@dataclasses.dataclass
class PostTrackRequest:
    request: Optional[Any] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})
    

@dataclasses.dataclass
class PostTrackResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    events_capture_response: Optional[shared_eventscaptureresponse.EventsCaptureResponse] = dataclasses.field(default=None)
    
