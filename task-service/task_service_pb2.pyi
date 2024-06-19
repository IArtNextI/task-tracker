from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateRequest(_message.Message):
    __slots__ = ["title", "content", "owner_id"]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    title: str
    content: str
    owner_id: str
    def __init__(self, title: _Optional[str] = ..., content: _Optional[str] = ..., owner_id: _Optional[str] = ...) -> None: ...

class CreateResponse(_message.Message):
    __slots__ = ["task_id"]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    def __init__(self, task_id: _Optional[str] = ...) -> None: ...

class UpdateRequest(_message.Message):
    __slots__ = ["task_id", "owner_id", "title", "content", "status"]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    owner_id: str
    title: str
    content: str
    status: str
    def __init__(self, task_id: _Optional[str] = ..., owner_id: _Optional[str] = ..., title: _Optional[str] = ..., content: _Optional[str] = ..., status: _Optional[str] = ...) -> None: ...

class UpdateResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class DeleteRequest(_message.Message):
    __slots__ = ["task_id", "owner_id"]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    owner_id: str
    def __init__(self, task_id: _Optional[str] = ..., owner_id: _Optional[str] = ...) -> None: ...

class DeleteResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetRequest(_message.Message):
    __slots__ = ["task_id"]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    def __init__(self, task_id: _Optional[str] = ...) -> None: ...

class GetResponse(_message.Message):
    __slots__ = ["title", "content", "status", "owner_id"]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    title: str
    content: str
    status: str
    owner_id: str
    def __init__(self, title: _Optional[str] = ..., content: _Optional[str] = ..., status: _Optional[str] = ..., owner_id: _Optional[str] = ...) -> None: ...

class GetListRequest(_message.Message):
    __slots__ = ["owner_id", "from_pos", "count"]
    OWNER_ID_FIELD_NUMBER: _ClassVar[int]
    FROM_POS_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    owner_id: str
    from_pos: int
    count: int
    def __init__(self, owner_id: _Optional[str] = ..., from_pos: _Optional[int] = ..., count: _Optional[int] = ...) -> None: ...

class GetListResponse(_message.Message):
    __slots__ = ["tasks"]
    TASKS_FIELD_NUMBER: _ClassVar[int]
    tasks: _containers.RepeatedCompositeFieldContainer[Task]
    def __init__(self, tasks: _Optional[_Iterable[_Union[Task, _Mapping]]] = ...) -> None: ...

class Task(_message.Message):
    __slots__ = ["task_id", "title", "content", "status"]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    title: str
    content: str
    status: str
    def __init__(self, task_id: _Optional[str] = ..., title: _Optional[str] = ..., content: _Optional[str] = ..., status: _Optional[str] = ...) -> None: ...
