"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import typing
import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class AddFileRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    FILE_NAME_FIELD_NUMBER: builtins.int
    HASH_FIELD_NUMBER: builtins.int
    URL_FIELD_NUMBER: builtins.int
    file_name: typing.Text
    hash: typing.Text
    url: typing.Text
    def __init__(self,
        *,
        file_name: typing.Text = ...,
        hash: typing.Text = ...,
        url: typing.Optional[typing.Text] = ...,
        ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["_url",b"_url","url",b"url"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["_url",b"_url","file_name",b"file_name","hash",b"hash","url",b"url"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["_url",b"_url"]) -> typing.Optional[typing_extensions.Literal["url"]]: ...
global___AddFileRequest = AddFileRequest

class CheckText(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    IS_GOOGLE_ON_FIELD_NUMBER: builtins.int
    DOCUMENT_FIELD_NUMBER: builtins.int
    USER_ID_FIELD_NUMBER: builtins.int
    is_google_on: builtins.bool
    document: typing.Text
    user_id: builtins.int
    def __init__(self,
        *,
        is_google_on: builtins.bool = ...,
        document: typing.Text = ...,
        user_id: builtins.int = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["document",b"document","is_google_on",b"is_google_on","user_id",b"user_id"]) -> None: ...
global___CheckText = CheckText