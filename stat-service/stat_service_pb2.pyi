from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetTop3LikedUsersResponse(_message.Message):
    __slots__ = ["user_ids"]
    class UserIdAndLikeCount(_message.Message):
        __slots__ = ["like_count", "user_id"]
        LIKE_COUNT_FIELD_NUMBER: _ClassVar[int]
        USER_ID_FIELD_NUMBER: _ClassVar[int]
        like_count: int
        user_id: str
        def __init__(self, user_id: _Optional[str] = ..., like_count: _Optional[int] = ...) -> None: ...
    USER_IDS_FIELD_NUMBER: _ClassVar[int]
    user_ids: _containers.RepeatedCompositeFieldContainer[GetTop3LikedUsersResponse.UserIdAndLikeCount]
    def __init__(self, user_ids: _Optional[_Iterable[_Union[GetTop3LikedUsersResponse.UserIdAndLikeCount, _Mapping]]] = ...) -> None: ...

class GetTop5LikedOrViewedTasksRequest(_message.Message):
    __slots__ = ["likes"]
    LIKES_FIELD_NUMBER: _ClassVar[int]
    likes: bool
    def __init__(self, likes: bool = ...) -> None: ...

class GetTop5LikedOrViewedTasksResponse(_message.Message):
    __slots__ = ["tasks"]
    class TaskIdAndCount(_message.Message):
        __slots__ = ["count", "task_id"]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        TASK_ID_FIELD_NUMBER: _ClassVar[int]
        count: int
        task_id: str
        def __init__(self, task_id: _Optional[str] = ..., count: _Optional[int] = ...) -> None: ...
    TASKS_FIELD_NUMBER: _ClassVar[int]
    tasks: _containers.RepeatedCompositeFieldContainer[GetTop5LikedOrViewedTasksResponse.TaskIdAndCount]
    def __init__(self, tasks: _Optional[_Iterable[_Union[GetTop5LikedOrViewedTasksResponse.TaskIdAndCount, _Mapping]]] = ...) -> None: ...

class GetViewsAndLikesRequest(_message.Message):
    __slots__ = ["task_id"]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    task_id: str
    def __init__(self, task_id: _Optional[str] = ...) -> None: ...

class GetViewsAndLikesResponse(_message.Message):
    __slots__ = ["likes", "task_id", "views"]
    LIKES_FIELD_NUMBER: _ClassVar[int]
    TASK_ID_FIELD_NUMBER: _ClassVar[int]
    VIEWS_FIELD_NUMBER: _ClassVar[int]
    likes: int
    task_id: str
    views: int
    def __init__(self, views: _Optional[int] = ..., likes: _Optional[int] = ..., task_id: _Optional[str] = ...) -> None: ...
