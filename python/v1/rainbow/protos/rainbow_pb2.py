# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: rainbow.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rrainbow.proto\x12\x1e\x63onvergedcomputing.org.grpc.v1\x1a\x19google/protobuf/any.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb3\x01\n\x07\x43ontent\x12\n\n\x02id\x18\x01 \x01(\t\x12\"\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x14.google.protobuf.Any\x12G\n\x08metadata\x18\x03 \x03(\x0b\x32\x35.convergedcomputing.org.grpc.v1.Content.MetadataEntry\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"m\n\x07Request\x12\x38\n\x07\x63ontent\x18\x01 \x01(\x0b\x32\'.convergedcomputing.org.grpc.v1.Content\x12(\n\x04sent\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"Y\n\x0fRegisterRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06secret\x18\x02 \x01(\t\x12(\n\x04sent\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x99\x01\n\x10SubmitJobRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07\x63luster\x18\x02 \x01(\t\x12\r\n\x05token\x18\x03 \x01(\t\x12\r\n\x05nodes\x18\x04 \x01(\x05\x12\r\n\x05tasks\x18\x05 \x01(\x05\x12\x0f\n\x07\x63ommand\x18\x06 \x01(\t\x12(\n\x04sent\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"p\n\x12RequestJobsRequest\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\x12\x0e\n\x06secret\x18\x02 \x01(\t\x12\x0f\n\x07maxJobs\x18\x03 \x01(\x05\x12(\n\x04sent\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"n\n\x11\x41\x63\x63\x65ptJobsRequest\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\x12\x0e\n\x06secret\x18\x02 \x01(\t\x12\x0e\n\x06jobids\x18\x03 \x03(\x05\x12(\n\x04sent\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xc8\x01\n\x08Response\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\x15\n\rmessage_count\x18\x02 \x01(\x03\x12\x1a\n\x12messages_processed\x18\x03 \x01(\x03\x12\x1a\n\x12processing_details\x18\x04 \x01(\t\"Y\n\nResultType\x12\x1b\n\x17RESULT_TYPE_UNSPECIFIED\x10\x00\x12\x17\n\x13RESULT_TYPE_SUCCESS\x10\x01\x12\x15\n\x11RESULT_TYPE_ERROR\x10\x02\"\x8e\x02\n\x10RegisterResponse\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\r\n\x05token\x18\x02 \x01(\t\x12\x0e\n\x06secret\x18\x03 \x01(\t\x12K\n\x06status\x18\x04 \x01(\x0e\x32;.convergedcomputing.org.grpc.v1.RegisterResponse.ResultType\"z\n\nResultType\x12\x18\n\x14REGISTER_UNSPECIFIED\x10\x00\x12\x14\n\x10REGISTER_SUCCESS\x10\x01\x12\x12\n\x0eREGISTER_ERROR\x10\x02\x12\x13\n\x0fREGISTER_DENIED\x10\x03\x12\x13\n\x0fREGISTER_EXISTS\x10\x04\"\xe3\x01\n\x11SubmitJobResponse\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\r\n\x05jobid\x18\x02 \x01(\x05\x12L\n\x06status\x18\x03 \x01(\x0e\x32<.convergedcomputing.org.grpc.v1.SubmitJobResponse.ResultType\"]\n\nResultType\x12\x16\n\x12SUBMIT_UNSPECIFIED\x10\x00\x12\x12\n\x0eSUBMIT_SUCCESS\x10\x01\x12\x10\n\x0cSUBMIT_ERROR\x10\x02\x12\x11\n\rSUBMIT_DENIED\x10\x03\"\xe8\x02\n\x13RequestJobsResponse\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12K\n\x04jobs\x18\x02 \x03(\x0b\x32=.convergedcomputing.org.grpc.v1.RequestJobsResponse.JobsEntry\x12N\n\x06status\x18\x03 \x01(\x0e\x32>.convergedcomputing.org.grpc.v1.RequestJobsResponse.ResultType\x1a+\n\tJobsEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"s\n\nResultType\x12\x1a\n\x16REQUEST_JOBS_NORESULTS\x10\x00\x12\x18\n\x14REQUEST_JOBS_SUCCESS\x10\x01\x12\x16\n\x12REQUEST_JOBS_ERROR\x10\x02\x12\x17\n\x13REQUEST_JOBS_DENIED\x10\x03\"\xd7\x01\n\x12\x41\x63\x63\x65ptJobsResponse\x12M\n\x06status\x18\x01 \x01(\x0e\x32=.convergedcomputing.org.grpc.v1.AcceptJobsResponse.ResultType\"r\n\nResultType\x12\x1b\n\x17RESULT_TYPE_UNSPECIFIED\x10\x00\x12\x17\n\x13RESULT_TYPE_PARTIAL\x10\x01\x12\x17\n\x13RESULT_TYPE_SUCCESS\x10\x02\x12\x15\n\x11RESULT_TYPE_ERROR\x10\x03\x32\x9e\x05\n\x10RainbowScheduler\x12m\n\x08Register\x12/.convergedcomputing.org.grpc.v1.RegisterRequest\x1a\x30.convergedcomputing.org.grpc.v1.RegisterResponse\x12p\n\tSubmitJob\x12\x30.convergedcomputing.org.grpc.v1.SubmitJobRequest\x1a\x31.convergedcomputing.org.grpc.v1.SubmitJobResponse\x12v\n\x0bRequestJobs\x12\x32.convergedcomputing.org.grpc.v1.RequestJobsRequest\x1a\x33.convergedcomputing.org.grpc.v1.RequestJobsResponse\x12s\n\nAcceptJobs\x12\x31.convergedcomputing.org.grpc.v1.AcceptJobsRequest\x1a\x32.convergedcomputing.org.grpc.v1.AcceptJobsResponse\x12[\n\x06Serial\x12\'.convergedcomputing.org.grpc.v1.Request\x1a(.convergedcomputing.org.grpc.v1.Response\x12_\n\x06Stream\x12\'.convergedcomputing.org.grpc.v1.Request\x1a(.convergedcomputing.org.grpc.v1.Response(\x01\x30\x01\x42\x33Z1github.com/converged-computing/rainbow/pkg/api/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'rainbow_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z1github.com/converged-computing/rainbow/pkg/api/v1'
  _globals['_CONTENT_METADATAENTRY']._options = None
  _globals['_CONTENT_METADATAENTRY']._serialized_options = b'8\001'
  _globals['_REQUESTJOBSRESPONSE_JOBSENTRY']._options = None
  _globals['_REQUESTJOBSRESPONSE_JOBSENTRY']._serialized_options = b'8\001'
  _globals['_CONTENT']._serialized_start=110
  _globals['_CONTENT']._serialized_end=289
  _globals['_CONTENT_METADATAENTRY']._serialized_start=242
  _globals['_CONTENT_METADATAENTRY']._serialized_end=289
  _globals['_REQUEST']._serialized_start=291
  _globals['_REQUEST']._serialized_end=400
  _globals['_REGISTERREQUEST']._serialized_start=402
  _globals['_REGISTERREQUEST']._serialized_end=491
  _globals['_SUBMITJOBREQUEST']._serialized_start=494
  _globals['_SUBMITJOBREQUEST']._serialized_end=647
  _globals['_REQUESTJOBSREQUEST']._serialized_start=649
  _globals['_REQUESTJOBSREQUEST']._serialized_end=761
  _globals['_ACCEPTJOBSREQUEST']._serialized_start=763
  _globals['_ACCEPTJOBSREQUEST']._serialized_end=873
  _globals['_RESPONSE']._serialized_start=876
  _globals['_RESPONSE']._serialized_end=1076
  _globals['_RESPONSE_RESULTTYPE']._serialized_start=987
  _globals['_RESPONSE_RESULTTYPE']._serialized_end=1076
  _globals['_REGISTERRESPONSE']._serialized_start=1079
  _globals['_REGISTERRESPONSE']._serialized_end=1349
  _globals['_REGISTERRESPONSE_RESULTTYPE']._serialized_start=1227
  _globals['_REGISTERRESPONSE_RESULTTYPE']._serialized_end=1349
  _globals['_SUBMITJOBRESPONSE']._serialized_start=1352
  _globals['_SUBMITJOBRESPONSE']._serialized_end=1579
  _globals['_SUBMITJOBRESPONSE_RESULTTYPE']._serialized_start=1486
  _globals['_SUBMITJOBRESPONSE_RESULTTYPE']._serialized_end=1579
  _globals['_REQUESTJOBSRESPONSE']._serialized_start=1582
  _globals['_REQUESTJOBSRESPONSE']._serialized_end=1942
  _globals['_REQUESTJOBSRESPONSE_JOBSENTRY']._serialized_start=1782
  _globals['_REQUESTJOBSRESPONSE_JOBSENTRY']._serialized_end=1825
  _globals['_REQUESTJOBSRESPONSE_RESULTTYPE']._serialized_start=1827
  _globals['_REQUESTJOBSRESPONSE_RESULTTYPE']._serialized_end=1942
  _globals['_ACCEPTJOBSRESPONSE']._serialized_start=1945
  _globals['_ACCEPTJOBSRESPONSE']._serialized_end=2160
  _globals['_ACCEPTJOBSRESPONSE_RESULTTYPE']._serialized_start=2046
  _globals['_ACCEPTJOBSRESPONSE_RESULTTYPE']._serialized_end=2160
  _globals['_RAINBOWSCHEDULER']._serialized_start=2163
  _globals['_RAINBOWSCHEDULER']._serialized_end=2833
# @@protoc_insertion_point(module_scope)