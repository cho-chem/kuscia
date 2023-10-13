# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kuscia/proto/api/v1alpha1/kusciaapi/health.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from kuscia.proto.api.v1alpha1 import common_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0kuscia/proto/api/v1alpha1/kusciaapi/health.proto\x12#kuscia.proto.api.v1alpha1.kusciaapi\x1a&kuscia/proto/api/v1alpha1/common.proto\"I\n\rHealthRequest\x12\x38\n\x06header\x18\x01 \x01(\x0b\x32(.kuscia.proto.api.v1alpha1.RequestHeader\"\x8a\x01\n\x0eHealthResponse\x12\x31\n\x06status\x18\x01 \x01(\x0b\x32!.kuscia.proto.api.v1alpha1.Status\x12\x45\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x37.kuscia.proto.api.v1alpha1.kusciaapi.HealthResponseData\"#\n\x12HealthResponseData\x12\r\n\x05ready\x18\x01 \x01(\x08\x32\x83\x01\n\rHealthService\x12r\n\x07healthZ\x12\x32.kuscia.proto.api.v1alpha1.kusciaapi.HealthRequest\x1a\x33.kuscia.proto.api.v1alpha1.kusciaapi.HealthResponseB^\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapib\x06proto3')



_HEALTHREQUEST = DESCRIPTOR.message_types_by_name['HealthRequest']
_HEALTHRESPONSE = DESCRIPTOR.message_types_by_name['HealthResponse']
_HEALTHRESPONSEDATA = DESCRIPTOR.message_types_by_name['HealthResponseData']
HealthRequest = _reflection.GeneratedProtocolMessageType('HealthRequest', (_message.Message,), {
  'DESCRIPTOR' : _HEALTHREQUEST,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.health_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.HealthRequest)
  })
_sym_db.RegisterMessage(HealthRequest)

HealthResponse = _reflection.GeneratedProtocolMessageType('HealthResponse', (_message.Message,), {
  'DESCRIPTOR' : _HEALTHRESPONSE,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.health_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.HealthResponse)
  })
_sym_db.RegisterMessage(HealthResponse)

HealthResponseData = _reflection.GeneratedProtocolMessageType('HealthResponseData', (_message.Message,), {
  'DESCRIPTOR' : _HEALTHRESPONSEDATA,
  '__module__' : 'kuscia.proto.api.v1alpha1.kusciaapi.health_pb2'
  # @@protoc_insertion_point(class_scope:kuscia.proto.api.v1alpha1.kusciaapi.HealthResponseData)
  })
_sym_db.RegisterMessage(HealthResponseData)

_HEALTHSERVICE = DESCRIPTOR.services_by_name['HealthService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n!org.secretflow.v1alpha1.kusciaapiZ9github.com/secretflow/kuscia/proto/api/v1alpha1/kusciaapi'
  _HEALTHREQUEST._serialized_start=129
  _HEALTHREQUEST._serialized_end=202
  _HEALTHRESPONSE._serialized_start=205
  _HEALTHRESPONSE._serialized_end=343
  _HEALTHRESPONSEDATA._serialized_start=345
  _HEALTHRESPONSEDATA._serialized_end=380
  _HEALTHSERVICE._serialized_start=383
  _HEALTHSERVICE._serialized_end=514
# @@protoc_insertion_point(module_scope)