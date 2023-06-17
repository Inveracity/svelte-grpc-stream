// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "proto/notifications/v1/notifications.proto" (package "notifications.v1", syntax proto3)
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
/**
 * @generated from protobuf message notifications.v1.Notification
 */
export interface Notification {
    /**
     * @generated from protobuf field: string subid = 1;
     */
    subid: string;
    /**
     * @generated from protobuf field: string text = 2;
     */
    text: string;
    /**
     * @generated from protobuf field: string sender = 3;
     */
    sender: string;
    /**
     * @generated from protobuf field: string recipient = 4;
     */
    recipient: string;
}
/**
 * @generated from protobuf message notifications.v1.NotificationServiceNotifyResponse
 */
export interface NotificationServiceNotifyResponse {
    /**
     * @generated from protobuf field: notifications.v1.Notification notifications = 1;
     */
    notifications?: Notification;
}
/**
 * @generated from protobuf message notifications.v1.SubscribeRequest
 */
export interface SubscribeRequest {
    /**
     * @generated from protobuf field: string subid = 1;
     */
    subid: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Notification$Type extends MessageType<Notification> {
    constructor() {
        super("notifications.v1.Notification", [
            { no: 1, name: "subid", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "text", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "sender", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "recipient", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Notification>): Notification {
        const message = { subid: "", text: "", sender: "", recipient: "" };
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
                case /* string subid */ 1:
                    message.subid = reader.string();
                    break;
                case /* string text */ 2:
                    message.text = reader.string();
                    break;
                case /* string sender */ 3:
                    message.sender = reader.string();
                    break;
                case /* string recipient */ 4:
                    message.recipient = reader.string();
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
        /* string subid = 1; */
        if (message.subid !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.subid);
        /* string text = 2; */
        if (message.text !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.text);
        /* string sender = 3; */
        if (message.sender !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.sender);
        /* string recipient = 4; */
        if (message.recipient !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.recipient);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message notifications.v1.Notification
 */
export const Notification = new Notification$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotificationServiceNotifyResponse$Type extends MessageType<NotificationServiceNotifyResponse> {
    constructor() {
        super("notifications.v1.NotificationServiceNotifyResponse", [
            { no: 1, name: "notifications", kind: "message", T: () => Notification }
        ]);
    }
    create(value?: PartialMessage<NotificationServiceNotifyResponse>): NotificationServiceNotifyResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<NotificationServiceNotifyResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: NotificationServiceNotifyResponse): NotificationServiceNotifyResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* notifications.v1.Notification notifications */ 1:
                    message.notifications = Notification.internalBinaryRead(reader, reader.uint32(), options, message.notifications);
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
    internalBinaryWrite(message: NotificationServiceNotifyResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* notifications.v1.Notification notifications = 1; */
        if (message.notifications)
            Notification.internalBinaryWrite(message.notifications, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message notifications.v1.NotificationServiceNotifyResponse
 */
export const NotificationServiceNotifyResponse = new NotificationServiceNotifyResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SubscribeRequest$Type extends MessageType<SubscribeRequest> {
    constructor() {
        super("notifications.v1.SubscribeRequest", [
            { no: 1, name: "subid", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SubscribeRequest>): SubscribeRequest {
        const message = { subid: "" };
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
                case /* string subid */ 1:
                    message.subid = reader.string();
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
        /* string subid = 1; */
        if (message.subid !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.subid);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message notifications.v1.SubscribeRequest
 */
export const SubscribeRequest = new SubscribeRequest$Type();
/**
 * @generated ServiceType for protobuf service notifications.v1.NotificationService
 */
export const NotificationService = new ServiceType("notifications.v1.NotificationService", [
    { name: "Notify", serverStreaming: true, options: {}, I: SubscribeRequest, O: NotificationServiceNotifyResponse }
]);
