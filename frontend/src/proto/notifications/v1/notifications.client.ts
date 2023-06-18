// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "proto/notifications/v1/notifications.proto" (package "notifications.v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { NotificationService } from "./notifications";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { NotificationServiceNotifyResponse } from "./notifications";
import type { SubscribeRequest } from "./notifications";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service notifications.v1.NotificationService
 */
export interface INotificationServiceClient {
    /**
     * @generated from protobuf rpc: Notify(notifications.v1.SubscribeRequest) returns (stream notifications.v1.NotificationServiceNotifyResponse);
     */
    notify(input: SubscribeRequest, options?: RpcOptions): ServerStreamingCall<SubscribeRequest, NotificationServiceNotifyResponse>;
}
/**
 * @generated from protobuf service notifications.v1.NotificationService
 */
export class NotificationServiceClient implements INotificationServiceClient, ServiceInfo {
    typeName = NotificationService.typeName;
    methods = NotificationService.methods;
    options = NotificationService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Notify(notifications.v1.SubscribeRequest) returns (stream notifications.v1.NotificationServiceNotifyResponse);
     */
    notify(input: SubscribeRequest, options?: RpcOptions): ServerStreamingCall<SubscribeRequest, NotificationServiceNotifyResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<SubscribeRequest, NotificationServiceNotifyResponse>("serverStreaming", this._transport, method, opt, input);
    }
}