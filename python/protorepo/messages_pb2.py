# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messages.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emessages.proto\x12\rkafka.message\"^\n\x0e\x41\x64\x64\x46ileRequest\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\x0c\n\x04hash\x18\x02 \x01(\t\x12\x10\n\x03url\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x11\n\tfull_path\x18\x04 \x01(\tB\x06\n\x04_url\"o\n\x0c\x41\x64\x64\x43hildFile\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\x0c\n\x04hash\x18\x02 \x01(\t\x12\x10\n\x03url\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x11\n\tfull_path\x18\x04 \x01(\t\x12\x11\n\tparent_id\x18\x05 \x01(\x03\x42\x06\n\x04_url\"D\n\tCheckText\x12\x14\n\x0cis_google_on\x18\x01 \x01(\x08\x12\x10\n\x08\x64ocument\x18\x02 \x01(\t\x12\x0f\n\x07user_id\x18\x03 \x01(\x03\"\xe9\x01\n\x0f\x43heckTextResult\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x17\n\x0f\x63heck_result_id\x18\x02 \x01(\x03\x12\x41\n\rcheck_results\x18\x03 \x01(\x0b\x32*.kafka.message.CheckTextResult.CheckResult\x1ai\n\x0b\x43heckResult\x12\x11\n\tfile_hash\x18\x01 \x01(\t\x12\x11\n\tfile_name\x18\x02 \x01(\t\x12\x1d\n\x15plagiarism_percentage\x18\x03 \x01(\x01\x12\x15\n\rextended_data\x18\x04 \x01(\tB\x0cZ\n./messagesb\x06proto3')



_ADDFILEREQUEST = DESCRIPTOR.message_types_by_name['AddFileRequest']
_ADDCHILDFILE = DESCRIPTOR.message_types_by_name['AddChildFile']
_CHECKTEXT = DESCRIPTOR.message_types_by_name['CheckText']
_CHECKTEXTRESULT = DESCRIPTOR.message_types_by_name['CheckTextResult']
_CHECKTEXTRESULT_CHECKRESULT = _CHECKTEXTRESULT.nested_types_by_name['CheckResult']
AddFileRequest = _reflection.GeneratedProtocolMessageType('AddFileRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDFILEREQUEST,
  '__module__' : 'messages_pb2'
  # @@protoc_insertion_point(class_scope:kafka.message.AddFileRequest)
  })
_sym_db.RegisterMessage(AddFileRequest)

AddChildFile = _reflection.GeneratedProtocolMessageType('AddChildFile', (_message.Message,), {
  'DESCRIPTOR' : _ADDCHILDFILE,
  '__module__' : 'messages_pb2'
  # @@protoc_insertion_point(class_scope:kafka.message.AddChildFile)
  })
_sym_db.RegisterMessage(AddChildFile)

CheckText = _reflection.GeneratedProtocolMessageType('CheckText', (_message.Message,), {
  'DESCRIPTOR' : _CHECKTEXT,
  '__module__' : 'messages_pb2'
  # @@protoc_insertion_point(class_scope:kafka.message.CheckText)
  })
_sym_db.RegisterMessage(CheckText)

CheckTextResult = _reflection.GeneratedProtocolMessageType('CheckTextResult', (_message.Message,), {

  'CheckResult' : _reflection.GeneratedProtocolMessageType('CheckResult', (_message.Message,), {
    'DESCRIPTOR' : _CHECKTEXTRESULT_CHECKRESULT,
    '__module__' : 'messages_pb2'
    # @@protoc_insertion_point(class_scope:kafka.message.CheckTextResult.CheckResult)
    })
  ,
  'DESCRIPTOR' : _CHECKTEXTRESULT,
  '__module__' : 'messages_pb2'
  # @@protoc_insertion_point(class_scope:kafka.message.CheckTextResult)
  })
_sym_db.RegisterMessage(CheckTextResult)
_sym_db.RegisterMessage(CheckTextResult.CheckResult)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\n./messages'
  _ADDFILEREQUEST._serialized_start=33
  _ADDFILEREQUEST._serialized_end=127
  _ADDCHILDFILE._serialized_start=129
  _ADDCHILDFILE._serialized_end=240
  _CHECKTEXT._serialized_start=242
  _CHECKTEXT._serialized_end=310
  _CHECKTEXTRESULT._serialized_start=313
  _CHECKTEXTRESULT._serialized_end=546
  _CHECKTEXTRESULT_CHECKRESULT._serialized_start=441
  _CHECKTEXTRESULT_CHECKRESULT._serialized_end=546
# @@protoc_insertion_point(module_scope)
