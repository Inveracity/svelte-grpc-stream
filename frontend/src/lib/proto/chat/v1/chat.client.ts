// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "proto/chat/v1/chat.proto" (package "proto.chat.v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { ChatService } from "./chat";
import type { SendResponse } from "./chat";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ChatMessage } from "./chat";
import type { ConnectRequest } from "./chat";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service proto.chat.v1.ChatService
 */
export interface IChatServiceClient {
    /**
     * @generated from protobuf rpc: Connect(proto.chat.v1.ConnectRequest) returns (stream proto.chat.v1.ChatMessage);
     */
    connect(input: ConnectRequest, options?: RpcOptions): ServerStreamingCall<ConnectRequest, ChatMessage>;
    /**
     * @generated from protobuf rpc: Send(proto.chat.v1.ChatMessage) returns (proto.chat.v1.SendResponse);
     */
    send(input: ChatMessage, options?: RpcOptions): UnaryCall<ChatMessage, SendResponse>;
}
/**
 * @generated from protobuf service proto.chat.v1.ChatService
 */
export class ChatServiceClient implements IChatServiceClient, ServiceInfo {
    typeName = ChatService.typeName;
    methods = ChatService.methods;
    options = ChatService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Connect(proto.chat.v1.ConnectRequest) returns (stream proto.chat.v1.ChatMessage);
     */
    connect(input: ConnectRequest, options?: RpcOptions): ServerStreamingCall<ConnectRequest, ChatMessage> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ConnectRequest, ChatMessage>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Send(proto.chat.v1.ChatMessage) returns (proto.chat.v1.SendResponse);
     */
    send(input: ChatMessage, options?: RpcOptions): UnaryCall<ChatMessage, SendResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<ChatMessage, SendResponse>("unary", this._transport, method, opt, input);
    }
}