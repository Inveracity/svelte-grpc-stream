// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "proto/notifications/v1/notifications.proto" (package "proto.notifications.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
// // RECEIVE NOTIFICATIONS ////

/**
 * The request payload to subscribe to a notification stream
 *
 * @generated from protobuf message proto.notifications.v1.SubscribeRequest
 */
export interface SubscribeRequest {
    /**
     * @generated from protobuf field: string channel_id = 1;
     */
    channelId: string;
    /**
     * @generated from protobuf field: string user_id = 2;
     */
    userId: string;
    /**
     * @generated from protobuf field: string last_ts = 3;
     */
    lastTs: string; // last timestamp received by the client
}
/**
 * Notification server stream after subscribing to a channel
 *
 * @generated from protobuf message proto.notifications.v1.Notification
 */
export interface Notification {
    /**
     * @generated from protobuf field: string channel_id = 1;
     */
    channelId: string;
    /**
     * @generated from protobuf field: string user_id = 2;
     */
    userId: string;
    /**
     * @generated from protobuf field: string text = 3;
     */
    text: string;
    /**
     * @generated from protobuf field: string ts = 4;
     */
    ts: string; // timestamp
}
// // SEND NOTIFICATIONS ////

/**
 * Streamed notification object
 *
 * @generated from protobuf message proto.notifications.v1.SendRequest
 */
export interface SendRequest {
    /**
     * @generated from protobuf field: string channel_id = 1;
     */
    channelId: string;
    /**
     * @generated from protobuf field: string user_id = 2;
     */
    userId: string;
    /**
     * @generated from protobuf field: string text = 3;
     */
    text: string;
}
/**
 * The response payload after sending a notification
 *
 * @generated from protobuf message proto.notifications.v1.SendResponse
 */
export interface SendResponse {
    /**
     * @generated from protobuf field: bool ok = 1;
     */
    ok: boolean;
    /**
     * @generated from protobuf field: string error = 2;
     */
    error: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class SubscribeRequest$Type extends MessageType<SubscribeRequest> {
    constructor() {
        super("proto.notifications.v1.SubscribeRequest", [
            { no: 1, name: "channel_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "user_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "last_ts", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SubscribeRequest>): SubscribeRequest {
        const message = { channelId: "", userId: "", lastTs: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SubscribeRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SubscribeRequest): SubscribeRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string channel_id */ 1:
                    message.channelId = reader.string();
                    break;
                case /* string user_id */ 2:
                    message.userId = reader.string();
                    break;
                case /* string last_ts */ 3:
                    message.lastTs = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SubscribeRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string channel_id = 1; */
        if (message.channelId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.channelId);
        /* string user_id = 2; */
        if (message.userId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.userId);
        /* string last_ts = 3; */
        if (message.lastTs !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.lastTs);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.notifications.v1.SubscribeRequest
 */
export const SubscribeRequest = new SubscribeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Notification$Type extends MessageType<Notification> {
    constructor() {
        super("proto.notifications.v1.Notification", [
            { no: 1, name: "channel_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "user_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "text", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "ts", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Notification>): Notification {
        const message = { channelId: "", userId: "", text: "", ts: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Notification>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Notification): Notification {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string channel_id */ 1:
                    message.channelId = reader.string();
                    break;
                case /* string user_id */ 2:
                    message.userId = reader.string();
                    break;
                case /* string text */ 3:
                    message.text = reader.string();
                    break;
                case /* string ts */ 4:
                    message.ts = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Notification, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string channel_id = 1; */
        if (message.channelId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.channelId);
        /* string user_id = 2; */
        if (message.userId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.userId);
        /* string text = 3; */
        if (message.text !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.text);
        /* string ts = 4; */
        if (message.ts !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.ts);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.notifications.v1.Notification
 */
export const Notification = new Notification$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SendRequest$Type extends MessageType<SendRequest> {
    constructor() {
        super("proto.notifications.v1.SendRequest", [
            { no: 1, name: "channel_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "user_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "text", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SendRequest>): SendRequest {
        const message = { channelId: "", userId: "", text: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SendRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SendRequest): SendRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string channel_id */ 1:
                    message.channelId = reader.string();
                    break;
                case /* string user_id */ 2:
                    message.userId = reader.string();
                    break;
                case /* string text */ 3:
                    message.text = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SendRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string channel_id = 1; */
        if (message.channelId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.channelId);
        /* string user_id = 2; */
        if (message.userId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.userId);
        /* string text = 3; */
        if (message.text !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.text);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.notifications.v1.SendRequest
 */
export const SendRequest = new SendRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SendResponse$Type extends MessageType<SendResponse> {
    constructor() {
        super("proto.notifications.v1.SendResponse", [
            { no: 1, name: "ok", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "error", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SendResponse>): SendResponse {
        const message = { ok: false, error: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SendResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SendResponse): SendResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool ok */ 1:
                    message.ok = reader.bool();
                    break;
                case /* string error */ 2:
                    message.error = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SendResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool ok = 1; */
        if (message.ok !== false)
            writer.tag(1, WireType.Varint).bool(message.ok);
        /* string error = 2; */
        if (message.error !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.error);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message proto.notifications.v1.SendResponse
 */
export const SendResponse = new SendResponse$Type();
/**
 * @generated ServiceType for protobuf service proto.notifications.v1.NotificationService
 */
export const NotificationService = new ServiceType("proto.notifications.v1.NotificationService", [
    { name: "Subscribe", serverStreaming: true, options: {}, I: SubscribeRequest, O: Notification },
    { name: "Send", options: {}, I: SendRequest, O: SendResponse }
]);
