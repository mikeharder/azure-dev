# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import models_pb2 as models__pb2
import user_config_pb2 as user__config__pb2

GRPC_GENERATED_VERSION = '1.71.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in user_config_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class UserConfigServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Get = channel.unary_unary(
                '/azdext.UserConfigService/Get',
                request_serializer=user__config__pb2.GetUserConfigRequest.SerializeToString,
                response_deserializer=user__config__pb2.GetUserConfigResponse.FromString,
                _registered_method=True)
        self.GetString = channel.unary_unary(
                '/azdext.UserConfigService/GetString',
                request_serializer=user__config__pb2.GetUserConfigStringRequest.SerializeToString,
                response_deserializer=user__config__pb2.GetUserConfigStringResponse.FromString,
                _registered_method=True)
        self.GetSection = channel.unary_unary(
                '/azdext.UserConfigService/GetSection',
                request_serializer=user__config__pb2.GetUserConfigSectionRequest.SerializeToString,
                response_deserializer=user__config__pb2.GetUserConfigSectionResponse.FromString,
                _registered_method=True)
        self.Set = channel.unary_unary(
                '/azdext.UserConfigService/Set',
                request_serializer=user__config__pb2.SetUserConfigRequest.SerializeToString,
                response_deserializer=models__pb2.EmptyResponse.FromString,
                _registered_method=True)
        self.Unset = channel.unary_unary(
                '/azdext.UserConfigService/Unset',
                request_serializer=user__config__pb2.UnsetUserConfigRequest.SerializeToString,
                response_deserializer=models__pb2.EmptyResponse.FromString,
                _registered_method=True)


class UserConfigServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Get(self, request, context):
        """Get retrieves a value by path
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetString(self, request, context):
        """GetString retrieves a value by path and returns it as a string
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetSection(self, request, context):
        """GetSection retrieves a section by path
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Set(self, request, context):
        """Set sets a value at a given path
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Unset(self, request, context):
        """Unset removes a value at a given path
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserConfigServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=user__config__pb2.GetUserConfigRequest.FromString,
                    response_serializer=user__config__pb2.GetUserConfigResponse.SerializeToString,
            ),
            'GetString': grpc.unary_unary_rpc_method_handler(
                    servicer.GetString,
                    request_deserializer=user__config__pb2.GetUserConfigStringRequest.FromString,
                    response_serializer=user__config__pb2.GetUserConfigStringResponse.SerializeToString,
            ),
            'GetSection': grpc.unary_unary_rpc_method_handler(
                    servicer.GetSection,
                    request_deserializer=user__config__pb2.GetUserConfigSectionRequest.FromString,
                    response_serializer=user__config__pb2.GetUserConfigSectionResponse.SerializeToString,
            ),
            'Set': grpc.unary_unary_rpc_method_handler(
                    servicer.Set,
                    request_deserializer=user__config__pb2.SetUserConfigRequest.FromString,
                    response_serializer=models__pb2.EmptyResponse.SerializeToString,
            ),
            'Unset': grpc.unary_unary_rpc_method_handler(
                    servicer.Unset,
                    request_deserializer=user__config__pb2.UnsetUserConfigRequest.FromString,
                    response_serializer=models__pb2.EmptyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'azdext.UserConfigService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('azdext.UserConfigService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class UserConfigService(object):
    """Missing associated documentation comment in .proto file."""

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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/azdext.UserConfigService/Get',
            user__config__pb2.GetUserConfigRequest.SerializeToString,
            user__config__pb2.GetUserConfigResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetString(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/azdext.UserConfigService/GetString',
            user__config__pb2.GetUserConfigStringRequest.SerializeToString,
            user__config__pb2.GetUserConfigStringResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetSection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/azdext.UserConfigService/GetSection',
            user__config__pb2.GetUserConfigSectionRequest.SerializeToString,
            user__config__pb2.GetUserConfigSectionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Set(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/azdext.UserConfigService/Set',
            user__config__pb2.SetUserConfigRequest.SerializeToString,
            models__pb2.EmptyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def Unset(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/azdext.UserConfigService/Unset',
            user__config__pb2.UnsetUserConfigRequest.SerializeToString,
            models__pb2.EmptyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
