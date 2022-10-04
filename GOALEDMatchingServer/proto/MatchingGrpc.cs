// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Matching.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981
#region Designer generated code

using grpc = global::Grpc.Core;

namespace MatchingService {
  public static partial class MatchingService
  {
    static readonly string __ServiceName = "MatchingService.MatchingService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.GetPlayerIdRequest> __Marshaller_MatchingService_GetPlayerIdRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.GetPlayerIdRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.GetPlayerIdResponse> __Marshaller_MatchingService_GetPlayerIdResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.GetPlayerIdResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.GetPublicRoomRequest> __Marshaller_MatchingService_GetPublicRoomRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.GetPublicRoomRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.GetPublicRoomResponse> __Marshaller_MatchingService_GetPublicRoomResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.GetPublicRoomResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.CreatePublicRoomRequest> __Marshaller_MatchingService_CreatePublicRoomRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.CreatePublicRoomRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.CreatePublicRoomResponse> __Marshaller_MatchingService_CreatePublicRoomResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.CreatePublicRoomResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.CreatePrivateRoomRequest> __Marshaller_MatchingService_CreatePrivateRoomRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.CreatePrivateRoomRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.CreatePrivateRoomResponse> __Marshaller_MatchingService_CreatePrivateRoomResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.CreatePrivateRoomResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.JoinPublicRoomRequest> __Marshaller_MatchingService_JoinPublicRoomRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.JoinPublicRoomRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.JoinPublicRoomResponse> __Marshaller_MatchingService_JoinPublicRoomResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.JoinPublicRoomResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.JoinPrivateRoomRequest> __Marshaller_MatchingService_JoinPrivateRoomRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.JoinPrivateRoomRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.JoinPrivateRoomResponse> __Marshaller_MatchingService_JoinPrivateRoomResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.JoinPrivateRoomResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.LeaveRoomRequest> __Marshaller_MatchingService_LeaveRoomRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.LeaveRoomRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MatchingService.LeaveRoomResponse> __Marshaller_MatchingService_LeaveRoomResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MatchingService.LeaveRoomResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.GetPlayerIdRequest, global::MatchingService.GetPlayerIdResponse> __Method_GetPlayerId = new grpc::Method<global::MatchingService.GetPlayerIdRequest, global::MatchingService.GetPlayerIdResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetPlayerId",
        __Marshaller_MatchingService_GetPlayerIdRequest,
        __Marshaller_MatchingService_GetPlayerIdResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.GetPublicRoomRequest, global::MatchingService.GetPublicRoomResponse> __Method_GetPublicRoom = new grpc::Method<global::MatchingService.GetPublicRoomRequest, global::MatchingService.GetPublicRoomResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetPublicRoom",
        __Marshaller_MatchingService_GetPublicRoomRequest,
        __Marshaller_MatchingService_GetPublicRoomResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.CreatePublicRoomRequest, global::MatchingService.CreatePublicRoomResponse> __Method_CreatePublicRoom = new grpc::Method<global::MatchingService.CreatePublicRoomRequest, global::MatchingService.CreatePublicRoomResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreatePublicRoom",
        __Marshaller_MatchingService_CreatePublicRoomRequest,
        __Marshaller_MatchingService_CreatePublicRoomResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.CreatePrivateRoomRequest, global::MatchingService.CreatePrivateRoomResponse> __Method_CreatePrivateRoom = new grpc::Method<global::MatchingService.CreatePrivateRoomRequest, global::MatchingService.CreatePrivateRoomResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreatePrivateRoom",
        __Marshaller_MatchingService_CreatePrivateRoomRequest,
        __Marshaller_MatchingService_CreatePrivateRoomResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.JoinPublicRoomRequest, global::MatchingService.JoinPublicRoomResponse> __Method_JoinPublicRoom = new grpc::Method<global::MatchingService.JoinPublicRoomRequest, global::MatchingService.JoinPublicRoomResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "JoinPublicRoom",
        __Marshaller_MatchingService_JoinPublicRoomRequest,
        __Marshaller_MatchingService_JoinPublicRoomResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.JoinPrivateRoomRequest, global::MatchingService.JoinPrivateRoomResponse> __Method_JoinPrivateRoom = new grpc::Method<global::MatchingService.JoinPrivateRoomRequest, global::MatchingService.JoinPrivateRoomResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "JoinPrivateRoom",
        __Marshaller_MatchingService_JoinPrivateRoomRequest,
        __Marshaller_MatchingService_JoinPrivateRoomResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MatchingService.LeaveRoomRequest, global::MatchingService.LeaveRoomResponse> __Method_LeaveRoom = new grpc::Method<global::MatchingService.LeaveRoomRequest, global::MatchingService.LeaveRoomResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "LeaveRoom",
        __Marshaller_MatchingService_LeaveRoomRequest,
        __Marshaller_MatchingService_LeaveRoomResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::MatchingService.MatchingReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of MatchingService</summary>
    [grpc::BindServiceMethod(typeof(MatchingService), "BindService")]
    public abstract partial class MatchingServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.GetPlayerIdResponse> GetPlayerId(global::MatchingService.GetPlayerIdRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.GetPublicRoomResponse> GetPublicRoom(global::MatchingService.GetPublicRoomRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.CreatePublicRoomResponse> CreatePublicRoom(global::MatchingService.CreatePublicRoomRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.CreatePrivateRoomResponse> CreatePrivateRoom(global::MatchingService.CreatePrivateRoomRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.JoinPublicRoomResponse> JoinPublicRoom(global::MatchingService.JoinPublicRoomRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.JoinPrivateRoomResponse> JoinPrivateRoom(global::MatchingService.JoinPrivateRoomRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MatchingService.LeaveRoomResponse> LeaveRoom(global::MatchingService.LeaveRoomRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for MatchingService</summary>
    public partial class MatchingServiceClient : grpc::ClientBase<MatchingServiceClient>
    {
      /// <summary>Creates a new client for MatchingService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchingServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for MatchingService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchingServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchingServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchingServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.GetPlayerIdResponse GetPlayerId(global::MatchingService.GetPlayerIdRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetPlayerId(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.GetPlayerIdResponse GetPlayerId(global::MatchingService.GetPlayerIdRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetPlayerId, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.GetPlayerIdResponse> GetPlayerIdAsync(global::MatchingService.GetPlayerIdRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetPlayerIdAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.GetPlayerIdResponse> GetPlayerIdAsync(global::MatchingService.GetPlayerIdRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetPlayerId, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.GetPublicRoomResponse GetPublicRoom(global::MatchingService.GetPublicRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetPublicRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.GetPublicRoomResponse GetPublicRoom(global::MatchingService.GetPublicRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetPublicRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.GetPublicRoomResponse> GetPublicRoomAsync(global::MatchingService.GetPublicRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetPublicRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.GetPublicRoomResponse> GetPublicRoomAsync(global::MatchingService.GetPublicRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetPublicRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.CreatePublicRoomResponse CreatePublicRoom(global::MatchingService.CreatePublicRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreatePublicRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.CreatePublicRoomResponse CreatePublicRoom(global::MatchingService.CreatePublicRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreatePublicRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.CreatePublicRoomResponse> CreatePublicRoomAsync(global::MatchingService.CreatePublicRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreatePublicRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.CreatePublicRoomResponse> CreatePublicRoomAsync(global::MatchingService.CreatePublicRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreatePublicRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.CreatePrivateRoomResponse CreatePrivateRoom(global::MatchingService.CreatePrivateRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreatePrivateRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.CreatePrivateRoomResponse CreatePrivateRoom(global::MatchingService.CreatePrivateRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreatePrivateRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.CreatePrivateRoomResponse> CreatePrivateRoomAsync(global::MatchingService.CreatePrivateRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreatePrivateRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.CreatePrivateRoomResponse> CreatePrivateRoomAsync(global::MatchingService.CreatePrivateRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreatePrivateRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.JoinPublicRoomResponse JoinPublicRoom(global::MatchingService.JoinPublicRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return JoinPublicRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.JoinPublicRoomResponse JoinPublicRoom(global::MatchingService.JoinPublicRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_JoinPublicRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.JoinPublicRoomResponse> JoinPublicRoomAsync(global::MatchingService.JoinPublicRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return JoinPublicRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.JoinPublicRoomResponse> JoinPublicRoomAsync(global::MatchingService.JoinPublicRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_JoinPublicRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.JoinPrivateRoomResponse JoinPrivateRoom(global::MatchingService.JoinPrivateRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return JoinPrivateRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.JoinPrivateRoomResponse JoinPrivateRoom(global::MatchingService.JoinPrivateRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_JoinPrivateRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.JoinPrivateRoomResponse> JoinPrivateRoomAsync(global::MatchingService.JoinPrivateRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return JoinPrivateRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.JoinPrivateRoomResponse> JoinPrivateRoomAsync(global::MatchingService.JoinPrivateRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_JoinPrivateRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.LeaveRoomResponse LeaveRoom(global::MatchingService.LeaveRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return LeaveRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MatchingService.LeaveRoomResponse LeaveRoom(global::MatchingService.LeaveRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_LeaveRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.LeaveRoomResponse> LeaveRoomAsync(global::MatchingService.LeaveRoomRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return LeaveRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MatchingService.LeaveRoomResponse> LeaveRoomAsync(global::MatchingService.LeaveRoomRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_LeaveRoom, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override MatchingServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new MatchingServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(MatchingServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_GetPlayerId, serviceImpl.GetPlayerId)
          .AddMethod(__Method_GetPublicRoom, serviceImpl.GetPublicRoom)
          .AddMethod(__Method_CreatePublicRoom, serviceImpl.CreatePublicRoom)
          .AddMethod(__Method_CreatePrivateRoom, serviceImpl.CreatePrivateRoom)
          .AddMethod(__Method_JoinPublicRoom, serviceImpl.JoinPublicRoom)
          .AddMethod(__Method_JoinPrivateRoom, serviceImpl.JoinPrivateRoom)
          .AddMethod(__Method_LeaveRoom, serviceImpl.LeaveRoom).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, MatchingServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_GetPlayerId, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.GetPlayerIdRequest, global::MatchingService.GetPlayerIdResponse>(serviceImpl.GetPlayerId));
      serviceBinder.AddMethod(__Method_GetPublicRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.GetPublicRoomRequest, global::MatchingService.GetPublicRoomResponse>(serviceImpl.GetPublicRoom));
      serviceBinder.AddMethod(__Method_CreatePublicRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.CreatePublicRoomRequest, global::MatchingService.CreatePublicRoomResponse>(serviceImpl.CreatePublicRoom));
      serviceBinder.AddMethod(__Method_CreatePrivateRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.CreatePrivateRoomRequest, global::MatchingService.CreatePrivateRoomResponse>(serviceImpl.CreatePrivateRoom));
      serviceBinder.AddMethod(__Method_JoinPublicRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.JoinPublicRoomRequest, global::MatchingService.JoinPublicRoomResponse>(serviceImpl.JoinPublicRoom));
      serviceBinder.AddMethod(__Method_JoinPrivateRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.JoinPrivateRoomRequest, global::MatchingService.JoinPrivateRoomResponse>(serviceImpl.JoinPrivateRoom));
      serviceBinder.AddMethod(__Method_LeaveRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MatchingService.LeaveRoomRequest, global::MatchingService.LeaveRoomResponse>(serviceImpl.LeaveRoom));
    }

  }
}
#endregion
