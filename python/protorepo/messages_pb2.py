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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emessages.proto\x12\rkafka.message\"Q\n\x0e\x41\x64\x64\x46ileRequest\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\x0c\n\x04hash\x18\x02 \x01(\t\x12\x0b\n\x03url\x18\x03 \x01(\t\x12\x11\n\tfull_path\x18\x04 \x01(\t\"E\n\tCheckText\x12\x14\n\x0cis_google_on\x18\x01 \x01(\x08\x12\x10\n\x08\x64ocument\x18\x02 \x01(\t\x12\x10\n\x08user_uid\x18\x04 \x01(\t\"Y\n\tCheckFile\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\x10\n\x08user_uid\x18\x02 \x01(\t\x12\x14\n\x0cis_google_on\x18\x03 \x01(\x08\x12\x11\n\tfull_path\x18\x04 \x01(\t\"\xd4\x01\n\x0f\x43heckTextResult\x12@\n\x0c\x63heck_result\x18\x03 \x01(\x0b\x32*.kafka.message.CheckTextResult.CheckResult\x12\x10\n\x08user_uid\x18\x04 \x01(\t\x1am\n\x0b\x43heckResult\x12\x11\n\tfile_hash\x18\x01 \x01(\t\x12\x11\n\tfile_name\x18\x02 \x01(\t\x12\x1d\n\x15plagiarism_percentage\x18\x03 \x01(\x01\x12\x19\n\x11\x65xtended_data_url\x18\x05 \x01(\tB\x0cZ\n./messagesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'messages_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\n./messages'
  _globals['_ADDFILEREQUEST']._serialized_start=33
  _globals['_ADDFILEREQUEST']._serialized_end=114
  _globals['_CHECKTEXT']._serialized_start=116
  _globals['_CHECKTEXT']._serialized_end=185
  _globals['_CHECKFILE']._serialized_start=187
  _globals['_CHECKFILE']._serialized_end=276
  _globals['_CHECKTEXTRESULT']._serialized_start=279
  _globals['_CHECKTEXTRESULT']._serialized_end=491
  _globals['_CHECKTEXTRESULT_CHECKRESULT']._serialized_start=382
  _globals['_CHECKTEXTRESULT_CHECKRESULT']._serialized_end=491
# @@protoc_insertion_point(module_scope)
