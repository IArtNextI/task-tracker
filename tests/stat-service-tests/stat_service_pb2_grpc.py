# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import stat_service_pb2 as stat__service__pb2


class StatServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetViewsAndLikes = channel.unary_unary(
                '/StatService/GetViewsAndLikes',
                request_serializer=stat__service__pb2.GetViewsAndLikesRequest.SerializeToString,
                response_deserializer=stat__service__pb2.GetViewsAndLikesResponse.FromString,
                )
        self.GetTop5LikedOrViewedTasks = channel.unary_unary(
                '/StatService/GetTop5LikedOrViewedTasks',
                request_serializer=stat__service__pb2.GetTop5LikedOrViewedTasksRequest.SerializeToString,
                response_deserializer=stat__service__pb2.GetTop5LikedOrViewedTasksResponse.FromString,
                )
        self.GetTop3LikedUsers = channel.unary_unary(
                '/StatService/GetTop3LikedUsers',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=stat__service__pb2.GetTop3LikedUsersResponse.FromString,
                )


class StatServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetViewsAndLikes(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTop5LikedOrViewedTasks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTop3LikedUsers(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_StatServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetViewsAndLikes': grpc.unary_unary_rpc_method_handler(
                    servicer.GetViewsAndLikes,
                    request_deserializer=stat__service__pb2.GetViewsAndLikesRequest.FromString,
                    response_serializer=stat__service__pb2.GetViewsAndLikesResponse.SerializeToString,
            ),
            'GetTop5LikedOrViewedTasks': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTop5LikedOrViewedTasks,
                    request_deserializer=stat__service__pb2.GetTop5LikedOrViewedTasksRequest.FromString,
                    response_serializer=stat__service__pb2.GetTop5LikedOrViewedTasksResponse.SerializeToString,
            ),
            'GetTop3LikedUsers': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTop3LikedUsers,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=stat__service__pb2.GetTop3LikedUsersResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'StatService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class StatService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetViewsAndLikes(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StatService/GetViewsAndLikes',
            stat__service__pb2.GetViewsAndLikesRequest.SerializeToString,
            stat__service__pb2.GetViewsAndLikesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetTop5LikedOrViewedTasks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StatService/GetTop5LikedOrViewedTasks',
            stat__service__pb2.GetTop5LikedOrViewedTasksRequest.SerializeToString,
            stat__service__pb2.GetTop5LikedOrViewedTasksResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetTop3LikedUsers(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StatService/GetTop3LikedUsers',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            stat__service__pb2.GetTop3LikedUsersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)