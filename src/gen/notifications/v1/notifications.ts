// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "notifications/v1/notifications.proto" (package "notifications.v1", syntax proto3)
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
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
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
 * @generated from protobuf message notifications.v1.NotifyRequest
 */
export interface NotifyRequest {
}
// @generated message type with reflection information, may provide speed optimized methods
class Notification$Type extends MessageType<Notification> {
    constructor() {
        super("notifications.v1.Notification", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Notification>): Notification {
        const message = { id: "", name: "" };
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
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
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
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
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
class NotifyRequest$Type extends MessageType<NotifyRequest> {
    constructor() {
        super("notifications.v1.NotifyRequest", []);
    }
    create(value?: PartialMessage<NotifyRequest>): NotifyRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<NotifyRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: NotifyRequest): NotifyRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: NotifyRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message notifications.v1.NotifyRequest
 */
export const NotifyRequest = new NotifyRequest$Type();
/**
 * @generated ServiceType for protobuf service notifications.v1.NotificationService
 */
export const NotificationService = new ServiceType("notifications.v1.NotificationService", [
    { name: "Notify", serverStreaming: true, options: {}, I: NotifyRequest, O: NotificationServiceNotifyResponse }
]);
