package me.figoxu.sso.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.14.0)",
    comments = "Source: api.proto")
public final class SsoServiceGrpc {

  private SsoServiceGrpc() {}

  public static final String SERVICE_NAME = "sso.SsoService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<Api.LoginInfoReq,
      Api.LoginInfoRsp> getGetLoginInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetLoginInfo",
      requestType = Api.LoginInfoReq.class,
      responseType = Api.LoginInfoRsp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<Api.LoginInfoReq,
      Api.LoginInfoRsp> getGetLoginInfoMethod() {
    io.grpc.MethodDescriptor<Api.LoginInfoReq, Api.LoginInfoRsp> getGetLoginInfoMethod;
    if ((getGetLoginInfoMethod = SsoServiceGrpc.getGetLoginInfoMethod) == null) {
      synchronized (SsoServiceGrpc.class) {
        if ((getGetLoginInfoMethod = SsoServiceGrpc.getGetLoginInfoMethod) == null) {
          SsoServiceGrpc.getGetLoginInfoMethod = getGetLoginInfoMethod = 
              io.grpc.MethodDescriptor.<Api.LoginInfoReq, Api.LoginInfoRsp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "sso.SsoService", "GetLoginInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  Api.LoginInfoReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  Api.LoginInfoRsp.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getGetLoginInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<Api.User,
      Api.User> getSaveUserInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SaveUserInfo",
      requestType = Api.User.class,
      responseType = Api.User.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<Api.User,
      Api.User> getSaveUserInfoMethod() {
    io.grpc.MethodDescriptor<Api.User, Api.User> getSaveUserInfoMethod;
    if ((getSaveUserInfoMethod = SsoServiceGrpc.getSaveUserInfoMethod) == null) {
      synchronized (SsoServiceGrpc.class) {
        if ((getSaveUserInfoMethod = SsoServiceGrpc.getSaveUserInfoMethod) == null) {
          SsoServiceGrpc.getSaveUserInfoMethod = getSaveUserInfoMethod = 
              io.grpc.MethodDescriptor.<Api.User, Api.User>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "sso.SsoService", "SaveUserInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  Api.User.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  Api.User.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getSaveUserInfoMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SsoServiceStub newStub(io.grpc.Channel channel) {
    return new SsoServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SsoServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new SsoServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SsoServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new SsoServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class SsoServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getLoginInfo(Api.LoginInfoReq request,
        io.grpc.stub.StreamObserver<Api.LoginInfoRsp> responseObserver) {
      asyncUnimplementedUnaryCall(getGetLoginInfoMethod(), responseObserver);
    }

    /**
     */
    public void saveUserInfo(Api.User request,
        io.grpc.stub.StreamObserver<Api.User> responseObserver) {
      asyncUnimplementedUnaryCall(getSaveUserInfoMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetLoginInfoMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                Api.LoginInfoReq,
                Api.LoginInfoRsp>(
                  this, METHODID_GET_LOGIN_INFO)))
          .addMethod(
            getSaveUserInfoMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                Api.User,
                Api.User>(
                  this, METHODID_SAVE_USER_INFO)))
          .build();
    }
  }

  /**
   */
  public static final class SsoServiceStub extends io.grpc.stub.AbstractStub<SsoServiceStub> {
    private SsoServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SsoServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SsoServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SsoServiceStub(channel, callOptions);
    }

    /**
     */
    public void getLoginInfo(Api.LoginInfoReq request,
        io.grpc.stub.StreamObserver<Api.LoginInfoRsp> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetLoginInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void saveUserInfo(Api.User request,
        io.grpc.stub.StreamObserver<Api.User> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSaveUserInfoMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class SsoServiceBlockingStub extends io.grpc.stub.AbstractStub<SsoServiceBlockingStub> {
    private SsoServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SsoServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SsoServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SsoServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public Api.LoginInfoRsp getLoginInfo(Api.LoginInfoReq request) {
      return blockingUnaryCall(
          getChannel(), getGetLoginInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public Api.User saveUserInfo(Api.User request) {
      return blockingUnaryCall(
          getChannel(), getSaveUserInfoMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class SsoServiceFutureStub extends io.grpc.stub.AbstractStub<SsoServiceFutureStub> {
    private SsoServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SsoServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SsoServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SsoServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<Api.LoginInfoRsp> getLoginInfo(
        Api.LoginInfoReq request) {
      return futureUnaryCall(
          getChannel().newCall(getGetLoginInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<Api.User> saveUserInfo(
        Api.User request) {
      return futureUnaryCall(
          getChannel().newCall(getSaveUserInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_LOGIN_INFO = 0;
  private static final int METHODID_SAVE_USER_INFO = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SsoServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SsoServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_LOGIN_INFO:
          serviceImpl.getLoginInfo((Api.LoginInfoReq) request,
              (io.grpc.stub.StreamObserver<Api.LoginInfoRsp>) responseObserver);
          break;
        case METHODID_SAVE_USER_INFO:
          serviceImpl.saveUserInfo((Api.User) request,
              (io.grpc.stub.StreamObserver<Api.User>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (SsoServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getGetLoginInfoMethod())
              .addMethod(getSaveUserInfoMethod())
              .build();
        }
      }
    }
    return result;
  }
}
