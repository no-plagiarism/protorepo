# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messages.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emessages.proto\x12\rkafka.message\"^\n\x0e\x41\x64\x64\x46ileRequest\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\x0c\n\x04hash\x18\x02 \x01(\t\x12\x10\n\x03url\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x11\n\tfull_path\x18\x04 \x01(\tB\x06\n\x04_url\"E\n\tCheckText\x12\x14\n\x0cis_google_on\x18\x01 \x01(\x08\x12\x10\n\x08\x64ocument\x18\x02 \x01(\t\x12\x10\n\x08user_uid\x18\x04 \x01(\t\"\xd0\x01\n\x0f\x43heckTextResult\x12@\n\x0c\x63heck_result\x18\x03 \x01(\x0b\x32*.kafka.message.CheckTextResult.CheckResult\x12\x10\n\x08user_uid\x18\x04 \x01(\t\x1ai\n\x0b\x43heckResult\x12\x11\n\tfile_hash\x18\x01 \x01(\t\x12\x11\n\tfile_name\x18\x02 \x01(\t\x12\x1d\n\x15plagiarism_percentage\x18\x03 \x01(\x01\x12\x15\n\rextended_data\x18\x04 \x01(\tB\x0cZ\n./messagesb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'messages_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\n./messages'
  _ADDFILEREQUEST._serialized_start=33
  _ADDFILEREQUEST._serialized_end=127
  _CHECKTEXT._serialized_start=129
  _CHECKTEXT._serialized_end=198
  _CHECKTEXTRESULT._serialized_start=201
  _CHECKTEXTRESULT._serialized_end=409
  _CHECKTEXTRESULT_CHECKRESULT._serialized_start=304
  _CHECKTEXTRESULT_CHECKRESULT._serialized_end=409
# @@protoc_insertion_point(module_scope)
