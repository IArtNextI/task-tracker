# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import task_service_pb2 as task__service__pb2


class TaskServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/TaskService/Create',
                request_serializer=task__service__pb2.CreateRequest.SerializeToString,
                response_deserializer=task__service__pb2.CreateResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/TaskService/Update',
                request_serializer=task__service__pb2.UpdateRequest.SerializeToString,
                response_deserializer=task__service__pb2.UpdateResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/TaskService/Delete',
                request_serializer=task__service__pb2.DeleteRequest.SerializeToString,
                response_deserializer=task__service__pb2.DeleteResponse.FromString,
                )
        self.Get = channel.unary_unary(
                '/TaskService/Get',
                request_serializer=task__service__pb2.GetRequest.SerializeToString,
                response_deserializer=task__service__pb2.GetResponse.FromString,
                )
        self.GetList = channel.unary_unary(
                '/TaskService/GetList',
                request_serializer=task__service__pb2.GetListRequest.SerializeToString,
                response_deserializer=task__service__pb2.GetListResponse.FromString,
                )


class TaskServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetList(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TaskServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=task__service__pb2.CreateRequest.FromString,
                    response_serializer=task__service__pb2.CreateResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=task__service__pb2.UpdateRequest.FromString,
                    response_serializer=task__service__pb2.UpdateResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=task__service__pb2.DeleteRequest.FromString,
                    response_serializer=task__service__pb2.DeleteResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=task__service__pb2.GetRequest.FromString,
                    response_serializer=task__service__pb2.GetResponse.SerializeToString,
            ),
            'GetList': grpc.unary_unary_rpc_method_handler(
                    servicer.GetList,
                    request_deserializer=task__service__pb2.GetListRequest.FromString,
                    response_serializer=task__service__pb2.GetListResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'TaskService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TaskService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TaskService/Create',
            task__service__pb2.CreateRequest.SerializeToString,
            task__service__pb2.CreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TaskService/Update',
            task__service__pb2.UpdateRequest.SerializeToString,
            task__service__pb2.UpdateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TaskService/Delete',
            task__service__pb2.DeleteRequest.SerializeToString,
            task__service__pb2.DeleteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TaskService/Get',
            task__service__pb2.GetRequest.SerializeToString,
            task__service__pb2.GetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TaskService/GetList',
            task__service__pb2.GetListRequest.SerializeToString,
            task__service__pb2.GetListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
