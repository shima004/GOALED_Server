// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: GameService.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981
#region Designer generated code

using grpc = global::Grpc.Core;

namespace GameService {
  public static partial class GameService
  {
    static readonly string __ServiceName = "GameService.GameService";

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
    static readonly grpc::Marshaller<global::GameService.Room> __Marshaller_GameService_Room = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.Room.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SyncPlayerDataRequest> __Marshaller_GameService_SyncPlayerDataRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SyncPlayerDataRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SyncPlayerDataResponse> __Marshaller_GameService_SyncPlayerDataResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SyncPlayerDataResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SendPlayerDataRequest> __Marshaller_GameService_SendPlayerDataRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SendPlayerDataRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SendPlayerDataResponse> __Marshaller_GameService_SendPlayerDataResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SendPlayerDataResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SyncObjectRequest> __Marshaller_GameService_SyncObjectRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SyncObjectRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SyncObjectResponse> __Marshaller_GameService_SyncObjectResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SyncObjectResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SendObjectRequest> __Marshaller_GameService_SendObjectRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SendObjectRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.SendObjectResponse> __Marshaller_GameService_SendObjectResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.SendObjectResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.CloseStreamRequest> __Marshaller_GameService_CloseStreamRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.CloseStreamRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GameService.CloseStreamResponse> __Marshaller_GameService_CloseStreamResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GameService.CloseStreamResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GameService.Room, global::GameService.Room> __Method_CreateRoom = new grpc::Method<global::GameService.Room, global::GameService.Room>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateRoom",
        __Marshaller_GameService_Room,
        __Marshaller_GameService_Room);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GameService.SyncPlayerDataRequest, global::GameService.SyncPlayerDataResponse> __Method_SyncPlayerData = new grpc::Method<global::GameService.SyncPlayerDataRequest, global::GameService.SyncPlayerDataResponse>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "SyncPlayerData",
        __Marshaller_GameService_SyncPlayerDataRequest,
        __Marshaller_GameService_SyncPlayerDataResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GameService.SendPlayerDataRequest, global::GameService.SendPlayerDataResponse> __Method_SendPlayerData = new grpc::Method<global::GameService.SendPlayerDataRequest, global::GameService.SendPlayerDataResponse>(
        grpc::MethodType.ClientStreaming,
        __ServiceName,
        "SendPlayerData",
        __Marshaller_GameService_SendPlayerDataRequest,
        __Marshaller_GameService_SendPlayerDataResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GameService.SyncObjectRequest, global::GameService.SyncObjectResponse> __Method_SyncObject = new grpc::Method<global::GameService.SyncObjectRequest, global::GameService.SyncObjectResponse>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "SyncObject",
        __Marshaller_GameService_SyncObjectRequest,
        __Marshaller_GameService_SyncObjectResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GameService.SendObjectRequest, global::GameService.SendObjectResponse> __Method_SendObject = new grpc::Method<global::GameService.SendObjectRequest, global::GameService.SendObjectResponse>(
        grpc::MethodType.ClientStreaming,
        __ServiceName,
        "SendObject",
        __Marshaller_GameService_SendObjectRequest,
        __Marshaller_GameService_SendObjectResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GameService.CloseStreamRequest, global::GameService.CloseStreamResponse> __Method_CloseStream = new grpc::Method<global::GameService.CloseStreamRequest, global::GameService.CloseStreamResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CloseStream",
        __Marshaller_GameService_CloseStreamRequest,
        __Marshaller_GameService_CloseStreamResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::GameService.GameServiceReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of GameService</summary>
    [grpc::BindServiceMethod(typeof(GameService), "BindService")]
    public abstract partial class GameServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::GameService.Room> CreateRoom(global::GameService.Room request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task SyncPlayerData(global::GameService.SyncPlayerDataRequest request, grpc::IServerStreamWriter<global::GameService.SyncPlayerDataResponse> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::GameService.SendPlayerDataResponse> SendPlayerData(grpc::IAsyncStreamReader<global::GameService.SendPlayerDataRequest> requestStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task SyncObject(global::GameService.SyncObjectRequest request, grpc::IServerStreamWriter<global::GameService.SyncObjectResponse> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::GameService.SendObjectResponse> SendObject(grpc::IAsyncStreamReader<global::GameService.SendObjectRequest> requestStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::GameService.CloseStreamResponse> CloseStream(global::GameService.CloseStreamRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for GameService</summary>
    public partial class GameServiceClient : grpc::ClientBase<GameServiceClient>
    {
      /// <summary>Creates a new client for GameService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public GameServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for GameService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public GameServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected GameServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected GameServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GameService.Room CreateRoom(global::GameService.Room request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GameService.Room CreateRoom(global::GameService.Room request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GameService.Room> CreateRoomAsync(global::GameService.Room request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GameService.Room> CreateRoomAsync(global::GameService.Room request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateRoom, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::GameService.SyncPlayerDataResponse> SyncPlayerData(global::GameService.SyncPlayerDataRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SyncPlayerData(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::GameService.SyncPlayerDataResponse> SyncPlayerData(global::GameService.SyncPlayerDataRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_SyncPlayerData, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncClientStreamingCall<global::GameService.SendPlayerDataRequest, global::GameService.SendPlayerDataResponse> SendPlayerData(grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SendPlayerData(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncClientStreamingCall<global::GameService.SendPlayerDataRequest, global::GameService.SendPlayerDataResponse> SendPlayerData(grpc::CallOptions options)
      {
        return CallInvoker.AsyncClientStreamingCall(__Method_SendPlayerData, null, options);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::GameService.SyncObjectResponse> SyncObject(global::GameService.SyncObjectRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SyncObject(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::GameService.SyncObjectResponse> SyncObject(global::GameService.SyncObjectRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_SyncObject, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncClientStreamingCall<global::GameService.SendObjectRequest, global::GameService.SendObjectResponse> SendObject(grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SendObject(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncClientStreamingCall<global::GameService.SendObjectRequest, global::GameService.SendObjectResponse> SendObject(grpc::CallOptions options)
      {
        return CallInvoker.AsyncClientStreamingCall(__Method_SendObject, null, options);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GameService.CloseStreamResponse CloseStream(global::GameService.CloseStreamRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CloseStream(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GameService.CloseStreamResponse CloseStream(global::GameService.CloseStreamRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CloseStream, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GameService.CloseStreamResponse> CloseStreamAsync(global::GameService.CloseStreamRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CloseStreamAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GameService.CloseStreamResponse> CloseStreamAsync(global::GameService.CloseStreamRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CloseStream, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override GameServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new GameServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(GameServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_CreateRoom, serviceImpl.CreateRoom)
          .AddMethod(__Method_SyncPlayerData, serviceImpl.SyncPlayerData)
          .AddMethod(__Method_SendPlayerData, serviceImpl.SendPlayerData)
          .AddMethod(__Method_SyncObject, serviceImpl.SyncObject)
          .AddMethod(__Method_SendObject, serviceImpl.SendObject)
          .AddMethod(__Method_CloseStream, serviceImpl.CloseStream).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, GameServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_CreateRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::GameService.Room, global::GameService.Room>(serviceImpl.CreateRoom));
      serviceBinder.AddMethod(__Method_SyncPlayerData, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::GameService.SyncPlayerDataRequest, global::GameService.SyncPlayerDataResponse>(serviceImpl.SyncPlayerData));
      serviceBinder.AddMethod(__Method_SendPlayerData, serviceImpl == null ? null : new grpc::ClientStreamingServerMethod<global::GameService.SendPlayerDataRequest, global::GameService.SendPlayerDataResponse>(serviceImpl.SendPlayerData));
      serviceBinder.AddMethod(__Method_SyncObject, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::GameService.SyncObjectRequest, global::GameService.SyncObjectResponse>(serviceImpl.SyncObject));
      serviceBinder.AddMethod(__Method_SendObject, serviceImpl == null ? null : new grpc::ClientStreamingServerMethod<global::GameService.SendObjectRequest, global::GameService.SendObjectResponse>(serviceImpl.SendObject));
      serviceBinder.AddMethod(__Method_CloseStream, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::GameService.CloseStreamRequest, global::GameService.CloseStreamResponse>(serviceImpl.CloseStream));
    }

  }
}
#endregion
