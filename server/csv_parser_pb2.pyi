from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class CSVParserRequest(_message.Message):
    __slots__ = ("columnName",)
    COLUMNNAME_FIELD_NUMBER: _ClassVar[int]
    columnName: str
    def __init__(self, columnName: _Optional[str] = ...) -> None: ...

class CSVParserResponse(_message.Message):
    __slots__ = ("rowCount",)
    ROWCOUNT_FIELD_NUMBER: _ClassVar[int]
    rowCount: int
    def __init__(self, rowCount: _Optional[int] = ...) -> None: ...
