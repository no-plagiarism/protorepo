"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import typing
import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class HashFile(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor
    NAME_FIELD_NUMBER: builtins.int
    URL_FIELD_NUMBER: builtins.int
    CONTENT_HASH_FIELD_NUMBER: builtins.int
    WORDS_FIELD_NUMBER: builtins.int
    name: typing.Text
    url: typing.Text
    content_hash: typing.Text
    @property
    def words(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[typing.Text]: ...
    def __init__(self,
        *,
        name: typing.Text = ...,
        url: typing.Text = ...,
        content_hash: typing.Text = ...,
        words: typing.Optional[typing.Iterable[typing.Text]] = ...,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["content_hash",b"content_hash","name",b"name","url",b"url","words",b"words"]) -> None: ...
global___HashFile = HashFile
