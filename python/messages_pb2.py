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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emessages.proto\x12\rkafka.message\"K\n\x0e\x41\x64\x64\x46ileRequest\x12\x11\n\tfile_name\x18\x01 \x01(\t\x12\x0c\n\x04hash\x18\x02 \x01(\t\x12\x10\n\x03url\x18\x03 \x01(\tH\x00\x88\x01\x01\x42\x06\n\x04_url\"D\n\tCheckText\x12\x14\n\x0cis_google_on\x18\x01 \x01(\x08\x12\x10\n\x08\x64ocument\x18\x02 \x01(\t\x12\x0f\n\x07user_id\x18\x03 \x01(\x03\x42\x0cZ\n./messagesb\x06proto3')



_ADDFILEREQUEST = DESCRIPTOR.message_types_by_name['AddFileRequest']
_CHECKTEXT = DESCRIPTOR.message_types_by_name['CheckText']
AddFileRequest = _reflection.GeneratedProtocolMessageType('AddFileRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDFILEREQUEST,
  '__module__' : 'messages_pb2'
  # @@protoc_insertion_point(class_scope:kafka.message.AddFileRequest)
  })
_sym_db.RegisterMessage(AddFileRequest)

CheckText = _reflection.GeneratedProtocolMessageType('CheckText', (_message.Message,), {
  'DESCRIPTOR' : _CHECKTEXT,
  '__module__' : 'messages_pb2'
  # @@protoc_insertion_point(class_scope:kafka.message.CheckText)
  })
_sym_db.RegisterMessage(CheckText)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\n./messages'
  _ADDFILEREQUEST._serialized_start=33
  _ADDFILEREQUEST._serialized_end=108
  _CHECKTEXT._serialized_start=110
  _CHECKTEXT._serialized_end=178
# @@protoc_insertion_point(module_scope)
