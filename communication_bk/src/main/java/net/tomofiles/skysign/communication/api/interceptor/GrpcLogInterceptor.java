package net.tomofiles.skysign.communication.api.interceptor;

import io.grpc.Metadata;
import io.grpc.MethodDescriptor;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptor;
import io.grpc.Status;
import org.lognet.springboot.grpc.GRpcGlobalInterceptor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// https://stackoverflow.com/questions/47155084/intercepting-logging-requests-and-responses-in-grpc
@GRpcGlobalInterceptor
public class GrpcLogInterceptor implements ServerInterceptor {
    private static final Logger logger = LoggerFactory.getLogger(GrpcLogInterceptor.class);

    @Override
    public <M, R> ServerCall.Listener<M> interceptCall(
            ServerCall<M, R> call,
            Metadata headers,
            ServerCallHandler<M, R> next) {
        GrpcServerCall<M, R> grpcServerCall = new GrpcServerCall<>(call);

        ServerCall.Listener<M> listener = next.startCall(grpcServerCall, headers);

        return new GrpcForwardingServerCallListener<M>(call.getMethodDescriptor(), listener) {
            @Override
            public void onMessage(M message) {
                logger.info("REQUEST , API: {}, Message: {}", methodName, message.toString().replaceAll("\\r\\n|\\r|\\n", " "));
                super.onMessage(message);
            }
        };
    }

    private class GrpcServerCall<M, R> extends ServerCall<M, R> {

        ServerCall<M, R> serverCall;

        protected GrpcServerCall(ServerCall<M, R> serverCall) {
            this.serverCall = serverCall;
        }

        @Override
        public void request(int numMessages) {
            serverCall.request(numMessages);
        }

        @Override
        public void sendHeaders(Metadata headers) {
            serverCall.sendHeaders(headers);
        }

        @Override
        public void sendMessage(R message) {
            logger.info("RESPONSE, API: {}, Message: {}", serverCall.getMethodDescriptor().getFullMethodName(), message.toString().replaceAll("\\r\\n|\\r|\\n", " "));
            serverCall.sendMessage(message);
        }

        @Override
        public void close(Status status, Metadata trailers) {
            serverCall.close(status, trailers);
        }

        @Override
        public boolean isCancelled() {
            return serverCall.isCancelled();
        }

        @Override
        public MethodDescriptor<M, R> getMethodDescriptor() {
            return serverCall.getMethodDescriptor();
        }
    }

    private class GrpcForwardingServerCallListener<M> extends io.grpc.ForwardingServerCallListener.SimpleForwardingServerCallListener<M> {

        String methodName;

        protected <R> GrpcForwardingServerCallListener(MethodDescriptor<M, R> method, ServerCall.Listener<M> listener) {
            super(listener);
            methodName = method.getFullMethodName();
        }
    }
}